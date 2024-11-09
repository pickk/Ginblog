import { createRouter, createWebHistory } from 'vue-router'

const ArticleList = () =>
    import(/* webpackChunkName: "group-index" */ '../components/ArticleList.vue')
const Detail = () =>
    import(/* webpackChunkName: "group-detail" */ '../components/Details.vue')
const Category = () =>
    import(/* webpackChunkName: "group-category" */ '../components/CateList.vue')
const Search = () =>
    import(/* webpackChunkName: "group-search" */ '../components/Search.vue')

// 定义路由表
const routes = [
  { path: '/', component: ArticleList, meta: { title: '欢迎来到GinBlog' } },
  {
    path: '/article/detail/:id',
    component: Detail,
    meta: { title: '动态标题' }, // 占位符
    props: true
  },
  {
    path: '/category/:cid',
    component: Category,
    meta: { title: '分类信息' },
    props: true
  },
  {
    path: '/search/:title',
    component: Search,
    meta: { title: '搜索结果' },
    props: true
  }
]

// 创建路由实例
const router = createRouter({
  history: createWebHistory(process.env.BASE_URL),
  routes
})

// 捕获路由 `push` 错误
const originalPush = router.push
router.push = function push(location) {
  return originalPush.call(this, location).catch(err => err)
}

// 路由守卫设置页面标题
router.beforeEach((to, from, next) => {
  if (to.path.startsWith('/article/detail')) {
    // 动态获取 sessionStorage 中的 title 用于该路由
    document.title = window.sessionStorage.getItem('title') || '加载中'
  } else if (to.meta.title) {
    document.title = to.meta.title
  } else {
    document.title = '加载中'
  }
  next()
})

export default router
