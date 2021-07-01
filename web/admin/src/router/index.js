import Vue from 'vue'
import VueRouter from 'vue-router'
import Login from '../views/Login.vue'
import Admin from '../views/Admin.vue'

// 页面路由组件
import Index from '../components/admin/Index'
import AddArt from '../components/article/AddArt'
import ArtList from '../components/article/ArtList'
import CateList from '../components/category/CateList'
import UserList from '../components/user/UserList'

Vue.use(VueRouter)

const routes = [
  {
    path: '/login',
    name: 'login',
    component: Login
  },
  {
    path: '/admin',
    name: 'admin',
    component: Admin,
    children: [
      {
        path: 'index',
        component: Index
      },
      {
        path: 'addart',
        component: AddArt
      },
      {
        path: 'addart/:id',
        component: AddArt,
        props: true
      },
      {
        path: 'artlist',
        component: ArtList
      },
      {
        path: 'catelist',
        component: CateList
      },
      {
        path: 'userlist',
        component: UserList
      }
    ]
  }
]

const router = new VueRouter({
  routes
})

router.beforeEach((to, from, next) => {
  const token = window.sessionStorage.getItem('token')
  if (to.path === '/login') return next()
  if (!token) return next('/login')
  next()
  // if (!token && to.path === '/admin') {
  //   next('/login')
  // } else {
  //   next()
  // }
})

export default router
