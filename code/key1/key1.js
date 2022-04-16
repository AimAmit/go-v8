(() => {
  var __getOwnPropNames = Object.getOwnPropertyNames;
  var __commonJS = (cb, mod) => function __require() {
    return mod || (0, cb[__getOwnPropNames(cb)[0]])((mod = { exports: {} }).exports, mod), mod.exports;
  };

  // math.js
  var require_math = __commonJS({
    "math.js"(exports, module) {
      var sum = (a, b) => a + b;
      module.exports = {
        sum
      };
    }
  });

  // index.js
  var require_key1 = __commonJS({
    "index.js"(exports, module) {
      var math = require_math();
      var main = (o) => {
        console.log(math.sum(10, 20));
        console.log(JSON.stringify(o));
        return 5;
      };
      module.exports = main;
    }
  });
  require_key1();
})();
