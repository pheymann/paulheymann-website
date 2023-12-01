---
layout: post
title: CDC Testing for MVPs (and Legacy Systems)
date:   2023-12-01
categories: engineering
head_title: CDC Testing for MVPs (and Legacy Systems)
meta_description: CDC Testing for MVPs (and Legacy Systems)
meta_keywords: software,engineering,testing,tdd,test-driven-development,cdc,consumer-driven-contract,pact
---

Writing tests isn’t particularly cheap so they fall down on my priority list when creating for example prototypes or MVPs. To some degree rightly so. I don’t know *What* I am doing yet, therefore, the *How* (my implementation) is changing regularly. Having to constantly add and adjust tests, a specification of something that is still in flow, feels like a drag at this point.

But eventually, it culminates in something tangible. The *What* and *How* become clearer and now I sit there with a bunch of untested code. The only way that I can know that it works is by manually testing it. Every. Time. This is the real drag now.

## Our MVP example
Okay, but how would we go about it? First, let’s add a small example. We build a prototype of a movie database. Think IMDB minus all the features. After some back and forth the first MVP looks like this:

<img class="gif" src="/assets/img/cdc-example-app.gif">

It won’t win a design award but it works. Under the hood, it is a basic React app (code [here](https://github.com/pheymann/paulheymann-website/tree/a3a713903954208e0cd691b133fa93c5e0da3ad1/examples/cdc/app)) talking to a Golang backend (code [here](https://github.com/pheymann/paulheymann-website/tree/a3a713903954208e0cd691b133fa93c5e0da3ad1/examples/cdc)). Of course, you can substitute that tech stack with whatever you want. What we will talk about in the following isn’t tied to a language or framework.

Now I could go in and add a whole bunch of unit and integration tests but at the same time, this MVP just shipped. People seem to use it (don’t ask me why). We have to fix bugs, add features, and grow the product before we run out of time, money, or both.

So, as it happens, we have to find another, less time-consuming path. Instead of focusing on each class and function why not look at it from a higher level? What happens when we test from the perspective of a user? I mean, at this moment, we are doing that anyway. Manually testing, or acting like a user, is the only tool that we have. So why not make that an automated process to speed up things?
## A simple user interaction test
What is it a user would see when they open our app?

```js
test("Case: User opens the app and sees multiple movies", async () => {
  await render(
    <MemoryRouter initialEntries={['/']}>
      <App  fetch={ mockFetch } />,
    </MemoryRouter>
  );

  expect(screen.getByText(/Harry Potter and the Sorcerer's Stone/i)).toBeInTheDocument();
  expect(screen.getByText(/Harry Potter and the Chamber of Secrets/i)).toBeInTheDocument();
  expect(screen.getByText(/Create Movie/i)).toBeInTheDocument();
}
```

Now, we could mock all requests via [mockFetch](https://github.com/pheymann/paulheymann-website/blob/a3a713903954208e0cd691b133fa93c5e0da3ad1/examples/cdc/app/src/cdc_tests/cdc.js#L44) but we would need to write additional, independent tests for the backend too. The appeal of manual end-to-end tests is that we know that the whole thing works when we are done testing. So instead of having separate sets of data, we put it all into a central file shared between the frontend and the backend.

```yaml
name: User opens the app and sees multiple movies
view: /
app:
  textShouldExist:
    - Harry Potter and the Sorcerer's Stone
    - Harry Potter and the Chamber of Secrets
callChain:
  - request:
      uri: /api/movie/all
      method: GET
      headers:
    response:
      type: success
      status: 200
      body:
        - id: "0"
          title: "Harry Potter and the Sorcerer's Stone"
          releaseYear: 2001
        - id: "1"
          title: "Harry Potter and the Chamber of Secrets"
          releaseYear: 2002
```

That is our contract between both systems and the representation of a user interaction.

Our test code becomes:

```js
test("Case: User opens the app and sees multiple movies", async () => {
  // helper function to load contracts from disk
  const contract = loadContract('home/no_movies_found.yaml');
  const mockFetch = new MockFetch(contract.callChain);

  await render(
    <MemoryRouter initialEntries={['/']}>
      <App  fetch={ mockFetch } />,
    </MemoryRouter>
  );

  app.textShouldExist.forEach((text, _) => {
    expect(screen.getByText(text)).toBeInTheDocument()
  });
  expect(screen.getByText(/Create Movie/i)).toBeInTheDocument();
}
```

To not repeat ourselves all the time I will extract all that contract loading and rendering into a separate function called [runContract](https://github.com/pheymann/paulheymann-website/blob/a3a713903954208e0cd691b133fa93c5e0da3ad1/examples/cdc/app/src/cdc_tests/cdc.js#L8). Now we end up with the following:

```js
runContract("home/multiple_movies_found.yaml", (app, _) => {
  expect(screen.getByText(/Create Movie/i)).toBeInTheDocument()

  app.textShouldExist.forEach((text, _) => {
    expect(screen.getByText(text)).toBeInTheDocument()
  });
});
```

Let’s take the same YAML and try to test the backend with it. We have a chain of request and response pairs. What we can do is simply run all requests through our business logic and see if it produces the expected outcome.

```go
func Test_CDC_HomeMultipleMoviesFound(t *testing.T) {
  contract := MustLoadCDC(cdcPath)

  t.Logf("Case: %s", contract.Name)

  movieServiceMock := &mockMovieService{movies: ???}

  for _, requestAndResponse := range contract.CallChain {
    request := httptest.NewRequest(
      requestAndResponse.Request.Method,
      requestAndResponse.Request.Uri,
      cdcutil.createBodyReader(requestAndResponse),
    )

    cdcutil.setHeaders(request, requestAndResponse)

    responseWriter := httptest.NewRecorder()

    ctx := listmovies.Context{Movies: movieServiceMock}
    ctx.Handle(w, r)

    response := responseWriter.Result()
    body, err := io.ReadAll(response.Body)
    assert.NoError(t, err)

    assert.Equal(t, requestAndResponse.Response.StatusCode, response.StatusCode)
    if requestAndResponse.Response.Type == "success" && requestAndResponse.Response.Body != nil {
      cdcutil.AssertEqualInterface(t, *requestAndResponse.Response.Body, body)
    } else if requestAndResponse.Response.Type == "error" {
      assert.Equal(
        t,
        requestAndResponse.Response.ErrorBody,
        string(body),
        "error body does not match"
      )
    }
  }
}
```

That is a mouthful. Let’s take it apart. After loading the contract file we create a mock for our movie database. Then we iterate through all the request and response pairs and evaluate if our code produces the same response as is stored in the YAML.

But wait. What do we put into the mock database? That is a piece of information we have to put into our contract too.

```yaml
name: User opens the app and sees multiple movies
view: /
database:
  movies:
    - id: "0"
      title: "Harry Potter and the Sorcerer's Stone"
      releaseYear: 2001
    - id: "1"
      title: "Harry Potter and the Chamber of Secrets"
      releaseYear: 2002
app:
  textShouldExist:
    - Harry Potter and the Sorcerer's Stone
    - Harry Potter and the Chamber of Secrets
callChain:
  - request:
      uri: /api/movie/all
      method: GET
    ...
```

Now we can modify our test code like this:

```go
func Test_CDC_HomeMultipleMoviesFound(t *testing.T) {
  contract := MustLoadCDC(cdcPath)

  t.Logf("Case: %s", contract.Name)

  movieServiceMock := &mockMovieService{movies: contract.Database.Movies}
  ...
```

And with a little cleaning up we get the following:

```go
func Test_CDC_HomeMultipleMoviesFound(t *testing.T) {
  RunContracts(
    t,
    "/home/multiple_movies_found.yaml",
    func(r *http.Request, w http.ResponseWriter, movieService movie.MovieService) {
      ctx := listmovies.Context{Movies: movieService}

      ctx.Handle(w, r)
    }
  )
}
```

All that utility code was extracted into [RunContracts](https://github.com/pheymann/paulheymann-website/blob/a3a713903954208e0cd691b133fa93c5e0da3ad1/examples/cdc/internal/cdc/runcontract.go#L18).

That is it. Now we can simply add more YAML contracts and add those small test cases to the app and backend. Of course, you can also test different endpoints in a single contract or for example, make sure authentication is honored when doing certain operations. Those contracts do not become more complicated. Have a look at a bigger example [here](https://github.com/pheymann/paulheymann-website/blob/a3a713903954208e0cd691b133fa93c5e0da3ad1/examples/cdc/cdc/create_movie/created.yaml).

And has all that a name? Yes, **Consumer Driven Contract** (short: CDC) **testing**.
## Legacy Systems
In my title, I also hinted at legacy systems. Sometimes we end up with that old piece of software. Nobody knows anymore how it works, there is no test coverage whatsoever, and to be frank the code is not in a shape that we can easily add unit tests. What can we do about it?

Writing a full test suite will take forever and getting that through might not be feasible. So what is the next best thing? Right, specify a bunch of interactions that you know the system needs to fulfill. Over time you can add all the edge cases you discover and through that grow your coverage. That is also one of the reasons why I like CDC testing for those scenarios. It neatly follows the Pareto principle. I get 80% of the return (test coverage) for 20% of the investment (engineering time) because I only define interaction contracts on a higher level and don’t tie my tests immediately to an implementation.
