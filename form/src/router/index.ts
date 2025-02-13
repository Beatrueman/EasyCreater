import { createRouter, createWebHashHistory } from 'vue-router'
import Login from '../views/LoginPage.vue'
import Regist from '../views/RegistPage.vue'
import Home from '../views/HomePage.vue'
import TemplatePage from '../views/TemplatePage.vue'
import MakePage from '../views/MakePage.vue'
import MakePageSecond from '../views/MakePageSecond.vue'
import AboutPage from '../views/AboutPage.vue'
import IndexPage from '../views/IndexPage.vue'
import MyResumePage from '../views/MyResumePage.vue'
import { jwtDecode } from 'jwt-decode'
import { defineAsyncComponent } from 'vue';
import { getTemplateData } from '../apis/api'

const routes = [{ path: '/', name: 'login', component: Login },
                { path: '/regist', name: 'regist', component: Regist },
                { path: '/home', name: 'home', component: Home, 
                  children: [{
                    path: 'template',
                    name: 'TemplatePage',
                    component: TemplatePage,
                  },
                  { path: 'template/first', 
                    name: 'MakePage', 
                    component: defineAsyncComponent(() => 
                      getTemplateData(1)
                    )
                  },
                  { path: 'template/second', 
                    name: 'MakePageSecond', 
                    component: defineAsyncComponent(() => 
                      getTemplateData(2)
                    )
                  },
                  { path: 'about', 
                    name: 'AboutPage', 
                    component: AboutPage
                  },
                  { path: 'index', 
                    name: 'IndexPage', 
                    component: IndexPage
                  },
                  { path: 'my_resume', 
                    name: 'MyResumePage', 
                    component: MyResumePage
                  },
                ]},

            ]

const router = createRouter({
  history: createWebHashHistory(import.meta.env.BASE_URL),
  routes,
})


router.beforeEach((to, from, next) => {
  if (to.name === 'login' || to.name === 'regist') {
    next()
    return 
  }

  const token = localStorage.getItem('jwt-token')
  if(!token) {
    next({ name: 'login' })
    return
  }
  try {
    const decoded: any = jwtDecode(token)
    const now = Date.now() / 1000 // 当前时间
    if (decoded.exp < now && decoded.exp ) {
      localStorage.removeItem('jwt-token')
      next({ name: 'login' })
      return
    }
  } catch(error) {
    console.error('JWT 解析失败:', error)
    localStorage.removeItem('jwt-token')
    next({ name: 'login' })
    return
  }

  next()
})

export default router
