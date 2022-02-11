# vue

## src/main.js
```
import Vue from 'vue'
import app from '@/app.vue'
import store from '@/store'
import router from '@/router'

Vue.config.productionTip = false

new Vue({
    store,
    router,
    render: h => h(app),
}).$mount('#app')
```

## src/app.vue
```
<template>
    <div id="app">
        <router-view></router-view>
    </div>
</template>
<script>
    export default {
        name: 'app'
    }
</script>

<style>

</style>
```

## src/store/index.js
```
import Vue from 'vue'
import Vuex from 'vuex'

Vue.use(Vuex)

export default new Vuex.Store({
    state: {},
    actions: {},
    modules: {},
    mutations: {},
})
```

## src/router/index.js
```
import Vue from 'vue'
import Router from 'vue-router'

Vue.use(Router)

const RouterPush = Router.prototype.push
Router.prototype.push = function push(to) {
    return RouterPush.call(this, to).catch(err => err)
}

const routes = [
    {
        path: '/',
        name: 'index',
        component: () => import('@/views/index'),
        meta: {title: '首页'},
    },
    {
        path: '/login',
        name: 'login',
        component: () => import('@/views/login'),
        meta: {title: '登录'},
    },
]

const router = new Router({
    mode: 'hash', //hash, history
    routes: routes,
})

router.beforeEach((to, from, next) => {
    if (to.meta.title) {
        document.title = to.meta.title
    }

    let token = sessionStorage.getItem('token')

    if (token) {
        if (to.path == '/login') {
            next({path: '/'})
        } else {
            next()
        }
    } else {
        if (to.path == '/login') {
            next()
        } else {
            next({path: '/login'})
        }
    }
})

export default router

```

## vue.config.js
```
'use strict'
const path = require('path')
const port = process.env.port || 8080
const resolve = dir => {
    return path.join(__dirname, dir)
}

module.exports = {
    publicPath: '/',
    outputDir: 'dist',
    assetsDir: '',
    indexPath: 'index.html',
    filenameHashing: true,
    lintOnSave: 'default',
    runtimeCompiler: false,
    transpileDependencies: [],
    productionSourceMap: true,
    integrity: false,
    configureWebpack: {},
    chainWebpack: config => {
        config.resolve.alias
            .set('@', resolve('src'))
            .set('style', resolve('src/style'))
            .set('assets', resolve('src/assets'))
            .set('components', resolve('src/components'))

        config.module
            .rule('svg')
            .exclude.add(resolve('src/static/icons/svg'))
            .end()

        config.module
            .rule('icons')
            .test(/\.svg$/)
            .include.add(resolve('src/static/icons/svg'))
            .end()
            .use('svg-sprite-loader')
            .loader('svg-sprite-loader')
            .options({symbolId: 'icon-[name]'})
            .end()
    },
    devServer: {
        port: port,
        open: false,
        proxy: {
            '/api': {
                target: 'http://localhost:9090',
                secure: false,
                changeOrigin: true,
                pathRewrite: {
                    '^/api': '',
                }
            }
        },
    },
    parallel: require('os').cpus().length > 1,
    pwa: {},
    pluginOptions: {
        electronBuilder: {
            nodeIntegration: true
        }
    },
}
```
