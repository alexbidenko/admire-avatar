import {createWebHistory, createRouter} from 'vue-router';
import AuthLayout from './layout/AuthLayout.vue';
import AuthPage from './pages/AuthPage.vue';
import RegistrationPage from './pages/RegistrationPage.vue';
import Main from '~/pages/MainPage.vue';
import GenerateImagePage from '~/pages/GenerateImagePage.vue';

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
        path: '/generateImages',
        name: 'generateImages',
        component: GenerateImagePage,
      },
    ],
  },
];

const router = createRouter({
  history: createWebHistory(),
  routes,
});

export default router;
