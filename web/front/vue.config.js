module.exports = {
  devServer: {
    host: '::', // 使用 IPv6 地址绑定服务器
    port: 8080,
  },
  transpileDependencies: ['vuetify'],
  assetsDir: 'static',
  chainWebpack: config => {
    config.plugin('html').tap(args => {
      args[0].title = '欢迎来到GinBlog'
      return args
    })
  },
  productionSourceMap: false
}
