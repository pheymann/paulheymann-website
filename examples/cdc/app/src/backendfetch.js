
export class BackendFetch {
  apply(uri, props) {
    return fetch(uri, props);
  }
}
