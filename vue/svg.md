# svg

## vue.config.js
```
const resolve = dir => {
    return path.join(__dirname, dir)
}

module.exports = {
    chainWebpack: config => {
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
}
```

## src/components/svg/index.js
```
import Vue from 'vue'
import SvgIcon from '@/components/svg/SvgIcon'

Vue.component('svg-icon', SvgIcon)

const requireAll = ctx => ctx.keys().map(ctx)
requireAll(require.context('@/static/icons/svg', false, /\.svg$/))
```

## src/components/svg/SvgIcon.vue
```
<template>
    <svg :class="svgClass" aria-hidden="true" v-on="$listeners">
        <use :xlink:href="iconName"></use>
    </svg>
</template>

<script>
    export default {
        name: "SvgIcon",
        props: {
            iconClass: {
                type: String,
                required: true,
            },
            className: {
                type: String,
                default: '',
            }
        },
        computed: {
            iconName() {
                return `#icon-${this.iconClass}`
            },
            svgClass() {
                if (this.className) {
                    return 'svg-icon' + this.className
                } else {
                    return 'svg-icon'
                }
            },
            styleExternalIcon() {
                return {
                    mask: `url(${this.iconClass}) no-repeat 50% 50%`,
                    '-webkit-mask': `url(${this.iconClass}) no-repeat 50% 50%`
                }
            },
        }
    }
</script>

<style scoped>
    .svg-icon {
        width: 1em;
        height: 1em;
        vertical-align: -0.15em;
        fill: currentColor;
        overflow: hidden;
    }

    .svg-external-icon {
        background-color: currentColor;
        mask-size: cover !important;
        display: inline-block;
    }
</style>
```

## src/main.js
```
import '@/components/svg'
```
