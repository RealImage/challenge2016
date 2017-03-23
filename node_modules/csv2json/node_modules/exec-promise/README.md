# exec-promise

[![Build Status](https://img.shields.io/travis/julien-f/nodejs-exec-promise/master.svg)](http://travis-ci.org/julien-f/nodejs-exec-promise)
[![Dependency Status](https://david-dm.org/julien-f/nodejs-exec-promise/status.svg?theme=shields.io)](https://david-dm.org/julien-f/nodejs-exec-promise)
[![devDependency Status](https://david-dm.org/julien-f/nodejs-exec-promise/dev-status.svg?theme=shields.io)](https://david-dm.org/julien-f/nodejs-exec-promise#info=devDependencies)

> Testable CLIs with promises

## Introduction

**TODO**

- executables should be testable
- the execution flow should be predictable and followable (promises)

## Install

Download [manually](https://github.com/julien-f/nodejs-exec-promise/releases) or with package-manager.

#### [npm](https://npmjs.org/package/exec-promise)

```
npm install --save exec-promise
```

This library requires promises support, for Node versions prior to 0.12 [see
this page](https://github.com/julien-f/js-promise-toolbox#usage) to
enable them.

## Example

### ES 2015

```javascript
import execPromise from 'exec-promise'

// - The command line arguments are passed as first parameter.
// - Node will exists as soon as the promise is settled (with a code
//   different than 0 in case of an error).
// - All errors are catched and properly displayed with a stack
//   trace.
// - Any returned value (i.e. not undefined) will be prettily
//   displayed
execPromise(async args => {
  // ... do what you want here!
})
```

### ES5

```javascript
module.exports = function (args) {
  if (args.indexOf('-h') !== -1) {
    return 'Usage: my-program [-h | -v]'
  }

  if (args.indexOf('-v') !== -1) {
    var pkg = require('./package')
    return 'MyProgram version ' + pkg.version
  }

  var server = require('http').createServer()
  server.listen(80)

  // The program will run until the server closes or encounters an
  // error.
  return require('event-to-promise')(server, 'close')
}

// Executes the exported function if this module has been called
// directly.
if (!module.parent) {
  require('exec-promise')(module.exports)
}
```

## Contributing

Contributions are *very* welcome, either on the documentation or on
the code.

You may:

- report any [issue](https://github.com/julien-f/human-format/issues)
  you've encountered;
- fork and create a pull request.

## License

ISC © [Julien Fontanet](http://julien.isonoe.net)
