import {createWebHistory, createRouter} from 'vue-router';
// import AuthLayout from './src/layout/AuthLayout.vue';
import AuthPage from './src/pages/AuthPage.vue';
import RegistrationPage from './src/pages/RegistrationPage.vue';

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
];

const router = createRouter({
  history: createWebHistory(),
  routes,
});

export default router;
