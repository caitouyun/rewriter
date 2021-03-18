const go = new Go();
WebAssembly.instantiateStreaming(fetch('main.wasm'), go.importObject).then(
  (result) => {
    console.log(result);
    go.run(result.instance);
  }
);
