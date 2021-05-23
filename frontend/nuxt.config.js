const colors = require('vuetify/es5/util/colors').default

module.exports = {
  ssr:false,
  server: {
    port: process.env.PORT || 3000,
    host: '0.0.0.0'
  },
  /*
   ** Headers of the page
   */
  head: {
    titleTemplate: '%s - ' + process.env.npm_package_name,
    title: process.env.npm_package_name || '',
    meta: [
      { charset: 'utf-8' },
      { name: 'viewport', content: 'width=device-width, initial-scale=1' },
      {
        hid: 'description',
        name: 'description',
        content: process.env.npm_package_description || ''
      }
    ],
    link: [{ rel: 'icon', type: 'image/x-icon', href: '/favicon.ico' }]
  },
  /*
   ** Customize the progress-bar color
   */
  loading: { color: '#fff' },
  /*
   ** Global CSS
   */
  css: [],
  /*
   ** Plugins to load before mounting the App
   */
  plugins: [],
  /*
   ** Nuxt.js dev-modules
   */
  buildModules: [
    // Doc: https://github.com/nuxt-community/eslint-module
    '@nuxtjs/eslint-module',
    '@nuxtjs/vuetify'
  ],
  /*
   ** Nuxt.js modules
   */
  modules: [
    // Doc: https://axios.nuxtjs.org/usage
    '@nuxtjs/axios',
    '@nuxtjs/auth-next',
    'cookie-universal-nuxt',
  ],
  router: {
    middleware: ['auth']
  },
  auth: {
    redirect: {
      login: '/login',
      logout: '/',
      callback: '/',
      home: '/scanner'
    },
    strategies: {
      local: false,
      keycloak: {
        scheme: 'oauth2',
        endpoints: {
          authorization:
            'https://csi.cmkl.ac.th/auth/realms/csi/protocol/openid-connect/auth',
          token:
            'https://csi.cmkl.ac.th/auth/realms/csi/protocol/openid-connect/token',
          userInfo:
            'https://csi.cmkl.ac.th/auth/realms/csi/protocol/openid-connect/userinfo',
          logout:
            'https://csi.cmkl.ac.th/auth/realms/csi/protocol/openid-connect/logout?redirect_uri=' +
            encodeURIComponent('https://csi.cmkl.ac.th')
        },
        token: {
          property: 'access_token',
          required: true,
          type: 'Bearer',
          name: 'Authorization',
          maxAge: 300
        },
        cookie:{

        },
        refreshToken: {
          property: 'refresh_token',
          maxAge: 60 * 60 * 24 * 30
        },
        responseType: 'code',
        grantType: 'authorization_code',
        clientId: 'frontend',
        scope: ['profile', 'email'],
        codeChallengeMethod: 'S256'
      }
    }
  },
  /*
   ** Axios module configuration
   ** See https://axios.nuxtjs.org/options
   */
  axios: {},
  /*
   ** vuetify module configuration
   ** https://github.com/nuxt-community/vuetify-module
   */
  vuetify: {
    customVariables: ['~/assets/variables.scss'],
    theme: {
      light: true,
      themes: {
        light: {
          primary: colors.blue.darken2,
          accent: colors.indigo.darken1,
          secondary: colors.amber.darken3,
          info: colors.teal.lighten1,
          warning: colors.amber.base,
          error: colors.deepOrange.accent4,
          success: colors.green.accent3
        }
      }
    }
  },
  /*
   ** Build configuration
   */
  build: {
    /*
     ** You can extend webpack config here
     */
    extend(config, ctx) {}
  },
  publicRuntimeConfig: {
    host: process.env.HOST
  }
}
