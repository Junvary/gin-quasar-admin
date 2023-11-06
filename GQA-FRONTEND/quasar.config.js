/* eslint-env node */

/*
 * This file runs in a Node context (it's NOT transpiled by Babel), so use only
 * the ES6 features that are supported by your Node version. https://node.green/
 */

// Configuration for your app
// https://v2.quasar.dev/quasar-cli-vite/quasar-config-js


const { configure } = require('quasar/wrappers');
const path = require('path');
const UnoCSS = require('unocss/vite').default;
const { presetUno } = require('unocss')

module.exports = configure(function (ctx) {
    return {


        // https://v2.quasar.dev/quasar-cli/prefetch-feature
        // preFetch: true,

        // app boot file (/src/boot)
        // --> boot files are part of "main.js"
        // https://v2.quasar.dev/quasar-cli/boot-files
        boot: [
            'set-default',
            'theme',
            'version',
            'i18n',
            'axios',
            'permission',
            'bus',
            'print',
            'components',
            'directive',
            'lang',
            'document',
        ],

        // https://v2.quasar.dev/quasar-cli-vite/quasar-config-js#css
        css: [
            'app.scss'
        ],

        // https://github.com/quasarframework/quasar/tree/dev/extras
        extras: [
            'material-icons', // optional, you are not bound to it
            'ionicons-v4',
            'mdi-v6',
            'fontawesome-v6',
            'eva-icons',
            'themify',
            'line-awesome',
            // 'roboto-font-latin-ext', // this or either 'roboto-font', NEVER both!
            'bootstrap-icons',
            'roboto-font', // optional, you are not bound to it

        ],

        // Full list of options: https://v2.quasar.dev/quasar-cli-vite/quasar-config-js#build
        build: {
            target: {
                browser: ['es2019', 'edge88', 'firefox78', 'chrome87', 'safari13.1'],
                node: 'node16'
            },

            vueRouterMode: 'hash', // available values: 'hash', 'history'
            // vueRouterBase,
            // vueDevtools,
            // vueOptionsAPI: false,

            // rebuildCache: true, // rebuilds Vite/linter/etc cache on startup

            // publicPath: '/',
            // analyze: true,
            env: {
                API: ctx.dev
                    // 测试代理地址
                    ? "http://127.0.0.1:8888/"
                    // 正式代理地址
                    : "http://81.68.159.232:8888/"
            },
            // rawDefine: {}
            // ignorePublicFolder: true,
            // minify: false,
            // polyfillModulePreload: true,
            // distDir

            extendViteConf(viteConf) {
                viteConf.plugins.push(
                    ...UnoCSS({
                        presets: [presetUno()],
                        safelist: ['bg-#545c64', 'bg-#000c17']
                    })
                )
            },
            // viteVuePluginOptions: {},

            vitePlugins: [
                ['@intlify/vite-plugin-vue-i18n', {
                    // if you want to use Vue I18n Legacy API, you need to set `compositionOnly: false`
                    // compositionOnly: false,

                    // if you want to use named tokens in your Vue I18n messages, such as 'Hello {name}',
                    // you need to set `runtimeOnly: false`
                    // runtimeOnly: false,

                    // you need to set i18n resource including paths !
                    include: path.resolve(__dirname, './src/i18n/**')
                }]
            ]
        },

        // Full list of options: https://v2.quasar.dev/quasar-cli-vite/quasar-config-js#devServer
        devServer: {
            // https: true
            open: true // opens browser window automatically
        },

        // https://v2.quasar.dev/quasar-cli-vite/quasar-config-js#framework
        framework: {
            config: {},

            iconSet: 'material-icons', // Quasar icon set
            lang: 'zh-CN', // Quasar language pack

            // For special cases outside of where the auto-import strategy can have an impact
            // (like functional components as one of the examples),
            // you can manually specify Quasar components/directives to be available everywhere:
            //
            // components: [],
            // directives: [],

            // Quasar plugins
            plugins: [
                'Notify',
                'Cookies',
                'SessionStorage',
                'LocalStorage',
                'Loading',
                'LoadingBar',
                'Dialog',
                'AppFullscreen',
            ]
        },

        // animations: 'all', // --- includes all animations
        // https://v2.quasar.dev/options/animations
        animations: 'all',

        // https://v2.quasar.dev/quasar-cli-vite/quasar-config-js#property-sourcefiles
        // sourceFiles: {
        //   rootComponent: 'src/App.vue',
        //   router: 'src/router/index',
        //   store: 'src/store/index',
        //   registerServiceWorker: 'src-pwa/register-service-worker',
        //   serviceWorker: 'src-pwa/custom-service-worker',
        //   pwaManifestFile: 'src-pwa/manifest.json',
        //   electronMain: 'src-electron/electron-main',
        //   electronPreload: 'src-electron/electron-preload'
        // },

        // https://v2.quasar.dev/quasar-cli/developing-ssr/configuring-ssr
        ssr: {
            // ssrPwaHtmlFilename: 'offline.html', // do NOT use index.html as name!
            // will mess up SSR

            // extendSSRWebserverConf (esbuildConf) {},
            // extendPackageJson (json) {},

            pwa: false,

            // manualStoreHydration: true,
            // manualPostHydrationTrigger: true,

            prodPort: 3000, // The default port that the production server should use
            // (gets superseded if process.env.PORT is specified at runtime)

            middlewares: [
                'render' // keep this as last one
            ]
        },

        // https://v2.quasar.dev/quasar-cli/developing-pwa/configuring-pwa
        pwa: {
            workboxMode: 'generateSW', // or 'injectManifest'
            injectPwaMetaTags: true,
            swFilename: 'sw.js',
            manifestFilename: 'manifest.json',
            useCredentialsForManifestTag: false,
            // useFilenameHashes: true,
            // extendGenerateSWOptions (cfg) {}
            // extendInjectManifestOptions (cfg) {},
            // extendManifestJson (json) {}
            // extendPWACustomSWConf (esbuildConf) {}
        },

        // Full list of options: https://v2.quasar.dev/quasar-cli/developing-cordova-apps/configuring-cordova
        cordova: {
            // noIosLegacyBuildFlag: true, // uncomment only if you know what you are doing
        },

        // Full list of options: https://v2.quasar.dev/quasar-cli/developing-capacitor-apps/configuring-capacitor
        capacitor: {
            hideSplashscreen: true
        },

        // Full list of options: https://v2.quasar.dev/quasar-cli/developing-electron-apps/configuring-electron
        electron: {
            // extendElectronMainConf (esbuildConf)
            // extendElectronPreloadConf (esbuildConf)

            inspectPort: 5858,

            bundler: 'packager', // 'packager' or 'builder'

            packager: {
                // https://github.com/electron-userland/electron-packager/blob/master/docs/api.md#options

                // OS X / Mac App Store
                // appBundleId: '',
                // appCategoryType: '',
                // osxSign: '',
                // protocol: 'myapp://path',

                // Windows only
                // win32metadata: { ... }
            },

            builder: {
                // https://www.electron.build/configuration/configuration

                appId: 'gqa-vite'
            }
        },

        // Full list of options: https://v2.quasar.dev/quasar-cli-vite/developing-browser-extensions/configuring-bex
        bex: {
            contentScripts: [
                'my-content-script'
            ],

            // extendBexScriptsConf (esbuildConf) {}
            // extendBexManifestJson (json) {}
        }
    }
});
