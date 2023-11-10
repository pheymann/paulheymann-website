import { render, waitFor } from '@testing-library/react';
import { MemoryRouter } from 'react-router-dom';
import yaml from 'js-yaml';
import fs from 'fs';
import { fail } from 'assert';
import App from './../App';

export function runContract(contractPath, assertFn, userInteractionFn = async () => {}) {
  const contract = loadContract(contractPath);
  const mockFetch = new MockFetch(contract.callChain);

  test("Case: " + contract.name, async () => {
    await render(
      <MemoryRouter initialEntries={[contract.view]}>
        <App  fetch={ mockFetch }
              token={ contract.app?.authorizer?.token }
        />,
      </MemoryRouter>
    );

    await userInteractionFn();

    await waitFor(() => assertFn(contract.app, mockFetch.requestAndResponseCallCounter));
  });
}

function loadContract(contractPath) {
  const contract = yaml.load(
    fs.readFileSync(`./../cdc/${contractPath}`, 'utf8')
  );

  // adjust header structure to align with fetch api
  contract.callChain.forEach(requestAndResponse => {
    const headerObj = {};
    requestAndResponse.request.headers?.forEach(pair => {
      headerObj[pair.name] = pair.value;
    });
    requestAndResponse.request.headerObj = headerObj;
  });

  return contract;
}

class MockFetch {

  constructor(requestsAndResponses) {
    this.requestsAndResponsesMap = new Map();
    this.requestAndResponseCallCounter = new Map();

    requestsAndResponses.forEach(requestAndResponse => {
      this.requestsAndResponsesMap.set(requestAndResponse.request.uri, requestAndResponse);
      this.requestAndResponseCallCounter.set(requestAndResponse.request.uri, 0);
    });
  }

  apply(uri, { method, headers, body }) {
    const requestAndResponse = this.requestsAndResponsesMap.get(uri);
    if (requestAndResponse === undefined) {
      fail(`No request and response found for uri: ${uri}\n${body}`);
    }

    if (!headers) {
      headers = {};
    }

    expect(method).toEqual(requestAndResponse.request.method);
    expect(headers).toEqual(requestAndResponse.request.headerObj);

    const callCounter = this.requestAndResponseCallCounter.get(uri);
    this.requestAndResponseCallCounter.set(uri, callCounter + 1);

    return Promise.resolve({
      status: requestAndResponse.response.status,
      text: () => JSON.stringify(requestAndResponse.response.body),
      json: () => requestAndResponse.response.body,
      ok: requestAndResponse.response.status < 300,
    });
  }
}
