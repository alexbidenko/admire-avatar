import {createWebHistory, createRouter} from 'vue-router';
import AuthLayout from './src/layout/AuthLayout.vue';
import AuthPage from './src/pages/AuthPage.vue';
import RegistrationPage from './src/pages/RegistrationPage.vue';
import Main from '~/pages/MainPage.vue';
import ListImages from './src/pages/ListImages.vue';

const routes = [
  {
    path: '/auth',
    name: 'Auth',
    component: AuthPage,
  },
  {
    path: '/sign-up',
    name: 'Sign-up',
    component: RegistrationPage,
  },
  {
    path: '/',
    name: 'Authorized',
    component: AuthLayout,
    children: [
      {
        path: '/',
        name: 'main',
        component: Main,
      },
      {
        path: '/listImages',
        name: 'listImages',
        component: ListImages,
      },
    ],
  },
];

const router = createRouter({
  history: createWebHistory(),
  routes,
});

export default router;
