'use strict';

define('cine-distributor/tests/adapters/application.jshint', ['exports'], function (exports) {
  'use strict';

  QUnit.module('JSHint | adapters/application.js');
  QUnit.test('should pass jshint', function (assert) {
    assert.expect(1);
    assert.ok(true, 'adapters/application.js should pass jshint.');
  });
});
define('cine-distributor/tests/adapters/distributor.jshint', ['exports'], function (exports) {
  'use strict';

  QUnit.module('JSHint | adapters/distributor.js');
  QUnit.test('should pass jshint', function (assert) {
    assert.expect(1);
    assert.ok(false, 'adapters/distributor.js should pass jshint.\nadapters/distributor.js: line 1, col 8, \'DS\' is defined but never used.\n\n1 error');
  });
});
define('cine-distributor/tests/adapters/place.jshint', ['exports'], function (exports) {
  'use strict';

  QUnit.module('JSHint | adapters/place.js');
  QUnit.test('should pass jshint', function (assert) {
    assert.expect(1);
    assert.ok(false, 'adapters/place.js should pass jshint.\nadapters/place.js: line 1, col 8, \'DS\' is defined but never used.\n\n1 error');
  });
});
define('cine-distributor/tests/app.jshint', ['exports'], function (exports) {
  'use strict';

  QUnit.module('JSHint | app.js');
  QUnit.test('should pass jshint', function (assert) {
    assert.expect(1);
    assert.ok(true, 'app.js should pass jshint.');
  });
});
define('cine-distributor/tests/components/flash-messages.jshint', ['exports'], function (exports) {
  'use strict';

  QUnit.module('JSHint | components/flash-messages.js');
  QUnit.test('should pass jshint', function (assert) {
    assert.expect(1);
    assert.ok(false, 'components/flash-messages.js should pass jshint.\ncomponents/flash-messages.js: line 3, col 3, Missing semicolon.\ncomponents/flash-messages.js: line 1, col 16, \'Ember\' is not defined.\ncomponents/flash-messages.js: line 2, col 18, \'Ember\' is not defined.\n\n3 errors');
  });
});
define('cine-distributor/tests/components/single-distributor.jshint', ['exports'], function (exports) {
  'use strict';

  QUnit.module('JSHint | components/single-distributor.js');
  QUnit.test('should pass jshint', function (assert) {
    assert.expect(1);
    assert.ok(false, 'components/single-distributor.js should pass jshint.\ncomponents/single-distributor.js: line 14, col 60, Expected \'!==\' and instead saw \'!=\'.\n\n1 error');
  });
});
define('cine-distributor/tests/helpers/destroy-app', ['exports', 'ember'], function (exports, _ember) {
  exports['default'] = destroyApp;

  function destroyApp(application) {
    _ember['default'].run(application, 'destroy');
  }
});
define('cine-distributor/tests/helpers/destroy-app.jshint', ['exports'], function (exports) {
  'use strict';

  QUnit.module('JSHint | helpers/destroy-app.js');
  QUnit.test('should pass jshint', function (assert) {
    assert.expect(1);
    assert.ok(true, 'helpers/destroy-app.js should pass jshint.');
  });
});
define('cine-distributor/tests/helpers/flash-message', ['exports', 'ember-cli-flash/flash/object'], function (exports, _emberCliFlashFlashObject) {

  _emberCliFlashFlashObject['default'].reopen({ init: function init() {} });
});
define('cine-distributor/tests/helpers/flash-message.jshint', ['exports'], function (exports) {
  'use strict';

  QUnit.module('JSHint | helpers/flash-message.js');
  QUnit.test('should pass jshint', function (assert) {
    assert.expect(1);
    assert.ok(true, 'helpers/flash-message.js should pass jshint.');
  });
});
define('cine-distributor/tests/helpers/if-cond.jshint', ['exports'], function (exports) {
  'use strict';

  QUnit.module('JSHint | helpers/if-cond.js');
  QUnit.test('should pass jshint', function (assert) {
    assert.expect(1);
    assert.ok(false, 'helpers/if-cond.js should pass jshint.\nhelpers/if-cond.js: line 10, col 26, Expected \'===\' and instead saw \'==\'.\nhelpers/if-cond.js: line 14, col 26, Expected \'!==\' and instead saw \'!=\'.\n\n2 errors');
  });
});
define('cine-distributor/tests/helpers/if-length.jshint', ['exports'], function (exports) {
  'use strict';

  QUnit.module('JSHint | helpers/if-length.js');
  QUnit.test('should pass jshint', function (assert) {
    assert.expect(1);
    assert.ok(false, 'helpers/if-length.js should pass jshint.\nhelpers/if-length.js: line 16, col 26, Expected \'===\' and instead saw \'==\'.\nhelpers/if-length.js: line 20, col 26, Expected \'!==\' and instead saw \'!=\'.\n\n2 errors');
  });
});
define('cine-distributor/tests/helpers/index-of-value.jshint', ['exports'], function (exports) {
  'use strict';

  QUnit.module('JSHint | helpers/index-of-value.js');
  QUnit.test('should pass jshint', function (assert) {
    assert.expect(1);
    assert.ok(true, 'helpers/index-of-value.js should pass jshint.');
  });
});
define('cine-distributor/tests/helpers/is-active.jshint', ['exports'], function (exports) {
  'use strict';

  QUnit.module('JSHint | helpers/is-active.js');
  QUnit.test('should pass jshint', function (assert) {
    assert.expect(1);
    assert.ok(true, 'helpers/is-active.js should pass jshint.');
  });
});
define('cine-distributor/tests/helpers/module-for-acceptance', ['exports', 'qunit', 'ember', 'cine-distributor/tests/helpers/start-app', 'cine-distributor/tests/helpers/destroy-app'], function (exports, _qunit, _ember, _cineDistributorTestsHelpersStartApp, _cineDistributorTestsHelpersDestroyApp) {
  var Promise = _ember['default'].RSVP.Promise;

  exports['default'] = function (name) {
    var options = arguments.length <= 1 || arguments[1] === undefined ? {} : arguments[1];

    (0, _qunit.module)(name, {
      beforeEach: function beforeEach() {
        this.application = (0, _cineDistributorTestsHelpersStartApp['default'])();

        if (options.beforeEach) {
          return options.beforeEach.apply(this, arguments);
        }
      },

      afterEach: function afterEach() {
        var _this = this;

        var afterEach = options.afterEach && options.afterEach.apply(this, arguments);
        return Promise.resolve(afterEach).then(function () {
          return (0, _cineDistributorTestsHelpersDestroyApp['default'])(_this.application);
        });
      }
    });
  };
});
define('cine-distributor/tests/helpers/module-for-acceptance.jshint', ['exports'], function (exports) {
  'use strict';

  QUnit.module('JSHint | helpers/module-for-acceptance.js');
  QUnit.test('should pass jshint', function (assert) {
    assert.expect(1);
    assert.ok(true, 'helpers/module-for-acceptance.js should pass jshint.');
  });
});
define('cine-distributor/tests/helpers/resolver', ['exports', 'cine-distributor/resolver', 'cine-distributor/config/environment'], function (exports, _cineDistributorResolver, _cineDistributorConfigEnvironment) {

  var resolver = _cineDistributorResolver['default'].create();

  resolver.namespace = {
    modulePrefix: _cineDistributorConfigEnvironment['default'].modulePrefix,
    podModulePrefix: _cineDistributorConfigEnvironment['default'].podModulePrefix
  };

  exports['default'] = resolver;
});
define('cine-distributor/tests/helpers/resolver.jshint', ['exports'], function (exports) {
  'use strict';

  QUnit.module('JSHint | helpers/resolver.js');
  QUnit.test('should pass jshint', function (assert) {
    assert.expect(1);
    assert.ok(true, 'helpers/resolver.js should pass jshint.');
  });
});
define('cine-distributor/tests/helpers/show-place.jshint', ['exports'], function (exports) {
  'use strict';

  QUnit.module('JSHint | helpers/show-place.js');
  QUnit.test('should pass jshint', function (assert) {
    assert.expect(1);
    assert.ok(false, 'helpers/show-place.js should pass jshint.\nhelpers/show-place.js: line 8, col 28, Missing semicolon.\nhelpers/show-place.js: line 9, col 23, Missing semicolon.\nhelpers/show-place.js: line 10, col 28, Expected \'===\' and instead saw \'==\'.\nhelpers/show-place.js: line 11, col 27, Missing semicolon.\nhelpers/show-place.js: line 13, col 33, Expected \'===\' and instead saw \'==\'.\nhelpers/show-place.js: line 14, col 31, Missing semicolon.\nhelpers/show-place.js: line 17, col 30, Missing semicolon.\nhelpers/show-place.js: line 20, col 231, Missing semicolon.\n\n8 errors');
  });
});
define('cine-distributor/tests/helpers/start-app', ['exports', 'ember', 'cine-distributor/app', 'cine-distributor/config/environment'], function (exports, _ember, _cineDistributorApp, _cineDistributorConfigEnvironment) {
  exports['default'] = startApp;

  function startApp(attrs) {
    var application = undefined;

    var attributes = _ember['default'].merge({}, _cineDistributorConfigEnvironment['default'].APP);
    attributes = _ember['default'].merge(attributes, attrs); // use defaults, but you can override;

    _ember['default'].run(function () {
      application = _cineDistributorApp['default'].create(attributes);
      application.setupForTesting();
      application.injectTestHelpers();
    });

    return application;
  }
});
define('cine-distributor/tests/helpers/start-app.jshint', ['exports'], function (exports) {
  'use strict';

  QUnit.module('JSHint | helpers/start-app.js');
  QUnit.test('should pass jshint', function (assert) {
    assert.expect(1);
    assert.ok(true, 'helpers/start-app.js should pass jshint.');
  });
});
define('cine-distributor/tests/initializers/component-router-injector.jshint', ['exports'], function (exports) {
  'use strict';

  QUnit.module('JSHint | initializers/component-router-injector.js');
  QUnit.test('should pass jshint', function (assert) {
    assert.expect(1);
    assert.ok(true, 'initializers/component-router-injector.js should pass jshint.');
  });
});
define('cine-distributor/tests/integration/components/list-filter-test', ['exports', 'ember-qunit'], function (exports, _emberQunit) {

  (0, _emberQunit.moduleForComponent)('list-filter', 'Integration | Component | list filter', {
    integration: true
  });

  (0, _emberQunit.test)('it renders', function (assert) {

    // Set any properties with this.set('myProperty', 'value');
    // Handle any actions with this.on('myAction', function(val) { ... });

    this.render(Ember.HTMLBars.template((function () {
      return {
        meta: {
          'revision': 'Ember@2.8.3',
          'loc': {
            'source': null,
            'start': {
              'line': 1,
              'column': 0
            },
            'end': {
              'line': 1,
              'column': 15
            }
          }
        },
        isEmpty: false,
        arity: 0,
        cachedFragment: null,
        hasRendered: false,
        buildFragment: function buildFragment(dom) {
          var el0 = dom.createDocumentFragment();
          var el1 = dom.createComment('');
          dom.appendChild(el0, el1);
          return el0;
        },
        buildRenderNodes: function buildRenderNodes(dom, fragment, contextualElement) {
          var morphs = new Array(1);
          morphs[0] = dom.createMorphAt(fragment, 0, 0, contextualElement);
          dom.insertBoundary(fragment, 0);
          dom.insertBoundary(fragment, null);
          return morphs;
        },
        statements: [['content', 'list-filter', ['loc', [null, [1, 0], [1, 15]]], 0, 0, 0, 0]],
        locals: [],
        templates: []
      };
    })()));

    assert.equal(this.$().text().trim(), '');

    // Template block usage:
    this.render(Ember.HTMLBars.template((function () {
      var child0 = (function () {
        return {
          meta: {
            'revision': 'Ember@2.8.3',
            'loc': {
              'source': null,
              'start': {
                'line': 2,
                'column': 4
              },
              'end': {
                'line': 4,
                'column': 4
              }
            }
          },
          isEmpty: false,
          arity: 0,
          cachedFragment: null,
          hasRendered: false,
          buildFragment: function buildFragment(dom) {
            var el0 = dom.createDocumentFragment();
            var el1 = dom.createTextNode('      template block text\n');
            dom.appendChild(el0, el1);
            return el0;
          },
          buildRenderNodes: function buildRenderNodes() {
            return [];
          },
          statements: [],
          locals: [],
          templates: []
        };
      })();

      return {
        meta: {
          'revision': 'Ember@2.8.3',
          'loc': {
            'source': null,
            'start': {
              'line': 1,
              'column': 0
            },
            'end': {
              'line': 5,
              'column': 2
            }
          }
        },
        isEmpty: false,
        arity: 0,
        cachedFragment: null,
        hasRendered: false,
        buildFragment: function buildFragment(dom) {
          var el0 = dom.createDocumentFragment();
          var el1 = dom.createTextNode('\n');
          dom.appendChild(el0, el1);
          var el1 = dom.createComment('');
          dom.appendChild(el0, el1);
          var el1 = dom.createTextNode('  ');
          dom.appendChild(el0, el1);
          return el0;
        },
        buildRenderNodes: function buildRenderNodes(dom, fragment, contextualElement) {
          var morphs = new Array(1);
          morphs[0] = dom.createMorphAt(fragment, 1, 1, contextualElement);
          return morphs;
        },
        statements: [['block', 'list-filter', [], [], 0, null, ['loc', [null, [2, 4], [4, 20]]]]],
        locals: [],
        templates: [child0]
      };
    })()));

    assert.equal(this.$().text().trim(), 'template block text');
  });
});
define('cine-distributor/tests/integration/components/list-filter-test.jshint', ['exports'], function (exports) {
  'use strict';

  QUnit.module('JSHint | integration/components/list-filter-test.js');
  QUnit.test('should pass jshint', function (assert) {
    assert.expect(1);
    assert.ok(true, 'integration/components/list-filter-test.js should pass jshint.');
  });
});
define('cine-distributor/tests/integration/components/single-distributor-test', ['exports', 'ember-qunit'], function (exports, _emberQunit) {

  (0, _emberQunit.moduleForComponent)('single-distributor', 'Integration | Component | single distributor', {
    integration: true
  });

  (0, _emberQunit.test)('it renders', function (assert) {

    // Set any properties with this.set('myProperty', 'value');
    // Handle any actions with this.on('myAction', function(val) { ... });

    this.render(Ember.HTMLBars.template((function () {
      return {
        meta: {
          'revision': 'Ember@2.8.3',
          'loc': {
            'source': null,
            'start': {
              'line': 1,
              'column': 0
            },
            'end': {
              'line': 1,
              'column': 22
            }
          }
        },
        isEmpty: false,
        arity: 0,
        cachedFragment: null,
        hasRendered: false,
        buildFragment: function buildFragment(dom) {
          var el0 = dom.createDocumentFragment();
          var el1 = dom.createComment('');
          dom.appendChild(el0, el1);
          return el0;
        },
        buildRenderNodes: function buildRenderNodes(dom, fragment, contextualElement) {
          var morphs = new Array(1);
          morphs[0] = dom.createMorphAt(fragment, 0, 0, contextualElement);
          dom.insertBoundary(fragment, 0);
          dom.insertBoundary(fragment, null);
          return morphs;
        },
        statements: [['content', 'single-distributor', ['loc', [null, [1, 0], [1, 22]]], 0, 0, 0, 0]],
        locals: [],
        templates: []
      };
    })()));

    assert.equal(this.$().text().trim(), '');

    // Template block usage:
    this.render(Ember.HTMLBars.template((function () {
      var child0 = (function () {
        return {
          meta: {
            'revision': 'Ember@2.8.3',
            'loc': {
              'source': null,
              'start': {
                'line': 2,
                'column': 4
              },
              'end': {
                'line': 4,
                'column': 4
              }
            }
          },
          isEmpty: false,
          arity: 0,
          cachedFragment: null,
          hasRendered: false,
          buildFragment: function buildFragment(dom) {
            var el0 = dom.createDocumentFragment();
            var el1 = dom.createTextNode('      template block text\n');
            dom.appendChild(el0, el1);
            return el0;
          },
          buildRenderNodes: function buildRenderNodes() {
            return [];
          },
          statements: [],
          locals: [],
          templates: []
        };
      })();

      return {
        meta: {
          'revision': 'Ember@2.8.3',
          'loc': {
            'source': null,
            'start': {
              'line': 1,
              'column': 0
            },
            'end': {
              'line': 5,
              'column': 2
            }
          }
        },
        isEmpty: false,
        arity: 0,
        cachedFragment: null,
        hasRendered: false,
        buildFragment: function buildFragment(dom) {
          var el0 = dom.createDocumentFragment();
          var el1 = dom.createTextNode('\n');
          dom.appendChild(el0, el1);
          var el1 = dom.createComment('');
          dom.appendChild(el0, el1);
          var el1 = dom.createTextNode('  ');
          dom.appendChild(el0, el1);
          return el0;
        },
        buildRenderNodes: function buildRenderNodes(dom, fragment, contextualElement) {
          var morphs = new Array(1);
          morphs[0] = dom.createMorphAt(fragment, 1, 1, contextualElement);
          return morphs;
        },
        statements: [['block', 'single-distributor', [], [], 0, null, ['loc', [null, [2, 4], [4, 27]]]]],
        locals: [],
        templates: [child0]
      };
    })()));

    assert.equal(this.$().text().trim(), 'template block text');
  });
});
define('cine-distributor/tests/integration/components/single-distributor-test.jshint', ['exports'], function (exports) {
  'use strict';

  QUnit.module('JSHint | integration/components/single-distributor-test.js');
  QUnit.test('should pass jshint', function (assert) {
    assert.expect(1);
    assert.ok(true, 'integration/components/single-distributor-test.js should pass jshint.');
  });
});
define('cine-distributor/tests/models/distributor.jshint', ['exports'], function (exports) {
  'use strict';

  QUnit.module('JSHint | models/distributor.js');
  QUnit.test('should pass jshint', function (assert) {
    assert.expect(1);
    assert.ok(true, 'models/distributor.js should pass jshint.');
  });
});
define('cine-distributor/tests/models/place.jshint', ['exports'], function (exports) {
  'use strict';

  QUnit.module('JSHint | models/place.js');
  QUnit.test('should pass jshint', function (assert) {
    assert.expect(1);
    assert.ok(true, 'models/place.js should pass jshint.');
  });
});
define('cine-distributor/tests/resolver.jshint', ['exports'], function (exports) {
  'use strict';

  QUnit.module('JSHint | resolver.js');
  QUnit.test('should pass jshint', function (assert) {
    assert.expect(1);
    assert.ok(true, 'resolver.js should pass jshint.');
  });
});
define('cine-distributor/tests/router.jshint', ['exports'], function (exports) {
  'use strict';

  QUnit.module('JSHint | router.js');
  QUnit.test('should pass jshint', function (assert) {
    assert.expect(1);
    assert.ok(true, 'router.js should pass jshint.');
  });
});
define('cine-distributor/tests/routes/application.jshint', ['exports'], function (exports) {
  'use strict';

  QUnit.module('JSHint | routes/application.js');
  QUnit.test('should pass jshint', function (assert) {
    assert.expect(1);
    assert.ok(false, 'routes/application.js should pass jshint.\nroutes/application.js: line 18, col 46, Expected \'===\' and instead saw \'==\'.\nroutes/application.js: line 19, col 68, Missing semicolon.\nroutes/application.js: line 37, col 93, Missing semicolon.\nroutes/application.js: line 44, col 55, Missing semicolon.\nroutes/application.js: line 45, col 46, Expected \'===\' and instead saw \'==\'.\nroutes/application.js: line 46, col 68, Missing semicolon.\nroutes/application.js: line 73, col 59, Missing semicolon.\nroutes/application.js: line 74, col 59, Missing semicolon.\nroutes/application.js: line 80, col 55, Expected \'===\' and instead saw \'==\'.\nroutes/application.js: line 87, col 35, \'reason\' is defined but never used.\nroutes/application.js: line 95, col 35, \'reason\' is defined but never used.\nroutes/application.js: line 77, col 29, \'place\' is defined but never used.\nroutes/application.js: line 59, col 13, \'jQuery\' is not defined.\nroutes/application.js: line 63, col 13, \'jQuery\' is not defined.\nroutes/application.js: line 75, col 13, \'jQuery\' is not defined.\nroutes/application.js: line 84, col 21, \'jQuery\' is not defined.\nroutes/application.js: line 85, col 25, \'jQuery\' is not defined.\nroutes/application.js: line 85, col 48, \'$\' is not defined.\n\n18 errors');
  });
});
define('cine-distributor/tests/routes/distributors.jshint', ['exports'], function (exports) {
  'use strict';

  QUnit.module('JSHint | routes/distributors.js');
  QUnit.test('should pass jshint', function (assert) {
    assert.expect(1);
    assert.ok(false, 'routes/distributors.js should pass jshint.\nroutes/distributors.js: line 11, col 76, Expected \'===\' and instead saw \'==\'.\nroutes/distributors.js: line 32, col 59, Expected \'===\' and instead saw \'==\'.\nroutes/distributors.js: line 35, col 15, Missing semicolon.\n\n3 errors');
  });
});
define('cine-distributor/tests/routes/place.jshint', ['exports'], function (exports) {
  'use strict';

  QUnit.module('JSHint | routes/place.js');
  QUnit.test('should pass jshint', function (assert) {
    assert.expect(1);
    assert.ok(true, 'routes/place.js should pass jshint.');
  });
});
define('cine-distributor/tests/serializers/application.jshint', ['exports'], function (exports) {
  'use strict';

  QUnit.module('JSHint | serializers/application.js');
  QUnit.test('should pass jshint', function (assert) {
    assert.expect(1);
    assert.ok(false, 'serializers/application.js should pass jshint.\nserializers/application.js: line 7, col 37, \'method\' is defined but never used.\nserializers/application.js: line 5, col 16, \'Ember\' is not defined.\nserializers/application.js: line 8, col 16, \'Ember\' is not defined.\nserializers/application.js: line 11, col 9, \'Ember\' is not defined.\n\n4 errors');
  });
});
define('cine-distributor/tests/serializers/distributor.jshint', ['exports'], function (exports) {
  'use strict';

  QUnit.module('JSHint | serializers/distributor.js');
  QUnit.test('should pass jshint', function (assert) {
    assert.expect(1);
    assert.ok(false, 'serializers/distributor.js should pass jshint.\nserializers/distributor.js: line 1, col 8, \'DS\' is defined but never used.\n\n1 error');
  });
});
define('cine-distributor/tests/serializers/place.jshint', ['exports'], function (exports) {
  'use strict';

  QUnit.module('JSHint | serializers/place.js');
  QUnit.test('should pass jshint', function (assert) {
    assert.expect(1);
    assert.ok(false, 'serializers/place.js should pass jshint.\nserializers/place.js: line 1, col 8, \'DS\' is defined but never used.\n\n1 error');
  });
});
define('cine-distributor/tests/test-helper', ['exports', 'cine-distributor/tests/helpers/resolver', 'cine-distributor/tests/helpers/flash-message', 'ember-qunit'], function (exports, _cineDistributorTestsHelpersResolver, _cineDistributorTestsHelpersFlashMessage, _emberQunit) {

  (0, _emberQunit.setResolver)(_cineDistributorTestsHelpersResolver['default']);
});
define('cine-distributor/tests/test-helper.jshint', ['exports'], function (exports) {
  'use strict';

  QUnit.module('JSHint | test-helper.js');
  QUnit.test('should pass jshint', function (assert) {
    assert.expect(1);
    assert.ok(true, 'test-helper.js should pass jshint.');
  });
});
define('cine-distributor/tests/transforms/array.jshint', ['exports'], function (exports) {
  'use strict';

  QUnit.module('JSHint | transforms/array.js');
  QUnit.test('should pass jshint', function (assert) {
    assert.expect(1);
    assert.ok(true, 'transforms/array.js should pass jshint.');
  });
});
define('cine-distributor/tests/unit/adapters/distributor-test', ['exports', 'ember-qunit'], function (exports, _emberQunit) {

  (0, _emberQunit.moduleFor)('adapter:distributor', 'Unit | Adapter | distributor', {
    // Specify the other units that are required for this test.
    // needs: ['serializer:foo']
  });

  // Replace this with your real tests.
  (0, _emberQunit.test)('it exists', function (assert) {
    var adapter = this.subject();
    assert.ok(adapter);
  });
});
define('cine-distributor/tests/unit/adapters/distributor-test.jshint', ['exports'], function (exports) {
  'use strict';

  QUnit.module('JSHint | unit/adapters/distributor-test.js');
  QUnit.test('should pass jshint', function (assert) {
    assert.expect(1);
    assert.ok(true, 'unit/adapters/distributor-test.js should pass jshint.');
  });
});
define('cine-distributor/tests/unit/adapters/place-test', ['exports', 'ember-qunit'], function (exports, _emberQunit) {

  (0, _emberQunit.moduleFor)('adapter:place', 'Unit | Adapter | place', {
    // Specify the other units that are required for this test.
    // needs: ['serializer:foo']
  });

  // Replace this with your real tests.
  (0, _emberQunit.test)('it exists', function (assert) {
    var adapter = this.subject();
    assert.ok(adapter);
  });
});
define('cine-distributor/tests/unit/adapters/place-test.jshint', ['exports'], function (exports) {
  'use strict';

  QUnit.module('JSHint | unit/adapters/place-test.js');
  QUnit.test('should pass jshint', function (assert) {
    assert.expect(1);
    assert.ok(true, 'unit/adapters/place-test.js should pass jshint.');
  });
});
define('cine-distributor/tests/unit/models/distributor-test', ['exports', 'ember-qunit'], function (exports, _emberQunit) {

  (0, _emberQunit.moduleForModel)('distributor', 'Unit | Model | distributor', {
    // Specify the other units that are required for this test.
    needs: []
  });

  (0, _emberQunit.test)('it exists', function (assert) {
    var model = this.subject();
    // let store = this.store();
    assert.ok(!!model);
  });
});
define('cine-distributor/tests/unit/models/distributor-test.jshint', ['exports'], function (exports) {
  'use strict';

  QUnit.module('JSHint | unit/models/distributor-test.js');
  QUnit.test('should pass jshint', function (assert) {
    assert.expect(1);
    assert.ok(true, 'unit/models/distributor-test.js should pass jshint.');
  });
});
define('cine-distributor/tests/unit/models/place-test', ['exports', 'ember-qunit'], function (exports, _emberQunit) {

  (0, _emberQunit.moduleForModel)('place', 'Unit | Model | place', {
    // Specify the other units that are required for this test.
    needs: []
  });

  (0, _emberQunit.test)('it exists', function (assert) {
    var model = this.subject();
    // let store = this.store();
    assert.ok(!!model);
  });
});
define('cine-distributor/tests/unit/models/place-test.jshint', ['exports'], function (exports) {
  'use strict';

  QUnit.module('JSHint | unit/models/place-test.js');
  QUnit.test('should pass jshint', function (assert) {
    assert.expect(1);
    assert.ok(true, 'unit/models/place-test.js should pass jshint.');
  });
});
define('cine-distributor/tests/unit/routes/distributor-test', ['exports', 'ember-qunit'], function (exports, _emberQunit) {

  (0, _emberQunit.moduleFor)('route:distributor', 'Unit | Route | distributor', {
    // Specify the other units that are required for this test.
    // needs: ['controller:foo']
  });

  (0, _emberQunit.test)('it exists', function (assert) {
    var route = this.subject();
    assert.ok(route);
  });
});
define('cine-distributor/tests/unit/routes/distributor-test.jshint', ['exports'], function (exports) {
  'use strict';

  QUnit.module('JSHint | unit/routes/distributor-test.js');
  QUnit.test('should pass jshint', function (assert) {
    assert.expect(1);
    assert.ok(true, 'unit/routes/distributor-test.js should pass jshint.');
  });
});
define('cine-distributor/tests/unit/routes/place-test', ['exports', 'ember-qunit'], function (exports, _emberQunit) {

  (0, _emberQunit.moduleFor)('route:place', 'Unit | Route | place', {
    // Specify the other units that are required for this test.
    // needs: ['controller:foo']
  });

  (0, _emberQunit.test)('it exists', function (assert) {
    var route = this.subject();
    assert.ok(route);
  });
});
define('cine-distributor/tests/unit/routes/place-test.jshint', ['exports'], function (exports) {
  'use strict';

  QUnit.module('JSHint | unit/routes/place-test.js');
  QUnit.test('should pass jshint', function (assert) {
    assert.expect(1);
    assert.ok(true, 'unit/routes/place-test.js should pass jshint.');
  });
});
define('cine-distributor/tests/unit/serializers/application-test', ['exports', 'ember-qunit'], function (exports, _emberQunit) {

  (0, _emberQunit.moduleForModel)('application', 'Unit | Serializer | application', {
    // Specify the other units that are required for this test.
    needs: ['serializer:application']
  });

  // Replace this with your real tests.
  (0, _emberQunit.test)('it serializes records', function (assert) {
    var record = this.subject();

    var serializedRecord = record.serialize();

    assert.ok(serializedRecord);
  });
});
define('cine-distributor/tests/unit/serializers/application-test.jshint', ['exports'], function (exports) {
  'use strict';

  QUnit.module('JSHint | unit/serializers/application-test.js');
  QUnit.test('should pass jshint', function (assert) {
    assert.expect(1);
    assert.ok(true, 'unit/serializers/application-test.js should pass jshint.');
  });
});
define('cine-distributor/tests/unit/serializers/distributor-test', ['exports', 'ember-qunit'], function (exports, _emberQunit) {

  (0, _emberQunit.moduleForModel)('distributor', 'Unit | Serializer | distributor', {
    // Specify the other units that are required for this test.
    needs: ['serializer:distributor']
  });

  // Replace this with your real tests.
  (0, _emberQunit.test)('it serializes records', function (assert) {
    var record = this.subject();

    var serializedRecord = record.serialize();

    assert.ok(serializedRecord);
  });
});
define('cine-distributor/tests/unit/serializers/distributor-test.jshint', ['exports'], function (exports) {
  'use strict';

  QUnit.module('JSHint | unit/serializers/distributor-test.js');
  QUnit.test('should pass jshint', function (assert) {
    assert.expect(1);
    assert.ok(true, 'unit/serializers/distributor-test.js should pass jshint.');
  });
});
define('cine-distributor/tests/unit/serializers/place-test', ['exports', 'ember-qunit'], function (exports, _emberQunit) {

  (0, _emberQunit.moduleForModel)('place', 'Unit | Serializer | place', {
    // Specify the other units that are required for this test.
    needs: ['serializer:place']
  });

  // Replace this with your real tests.
  (0, _emberQunit.test)('it serializes records', function (assert) {
    var record = this.subject();

    var serializedRecord = record.serialize();

    assert.ok(serializedRecord);
  });
});
define('cine-distributor/tests/unit/serializers/place-test.jshint', ['exports'], function (exports) {
  'use strict';

  QUnit.module('JSHint | unit/serializers/place-test.js');
  QUnit.test('should pass jshint', function (assert) {
    assert.expect(1);
    assert.ok(true, 'unit/serializers/place-test.js should pass jshint.');
  });
});
/* jshint ignore:start */

require('cine-distributor/tests/test-helper');
EmberENV.TESTS_FILE_LOADED = true;

/* jshint ignore:end */
//# sourceMappingURL=tests.map
