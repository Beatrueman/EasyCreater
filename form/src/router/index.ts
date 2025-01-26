import { createRouter, createWebHashHistory } from 'vue-router'
import Login from '../views/LoginPage.vue'
import Regist from '../views/RegistPage.vue'
import Home from '../views/HomePage.vue'
import AboutPage from '../views/AboutPage.vue'
import MakePage from '../views/MakePage.vue'

const routes = [{ path: '/', name: 'login', component: Login },
                { path: '/regist', name: 'regist', component: Regist },
                { path: '/home', name: 'home', component: Home, 
                  children: [{
                    path: 'about',
                    name: 'AboutPage',
                    component: AboutPage,
                  },
                  { path: 'make', 
                    name: 'MakePage', 
                    component: MakePage 
                  }
                ]},

            ]

const router = createRouter({
  history: createWebHashHistory(import.meta.env.BASE_URL),
  routes,
})

router.beforeEach((to, from, next) => {
  if (to.name === 'login' || to.name === 'regist') {
    next()
  } else if (to.name !== 'login' && !localStorage.getItem('jwt-token')) {
    // 如果没有登录，且不是登录或注册页面，重定向到登录页面
    next({ name: 'login' })
  } else {
    // 如果已登录或目标路由不需要登录，继续访问
    next()
  }
})

export default router
