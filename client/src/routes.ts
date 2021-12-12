import {createWebHistory, createRouter} from 'vue-router';
import AuthLayout from './layout/AuthLayout.vue';
import AuthPage from './pages/AuthPage.vue';
import RegistrationPage from './pages/RegistrationPage.vue';
import MainPage from '~/pages/MainPage.vue';
import FolderPage from '~/pages/FolderPage.vue';
import GeneratePage from '~/pages/GeneratePage.vue';
import PasswordPage from '~/pages/PasswordPage.vue';
import PrintsPage from '~/pages/PrintsPage.vue';
import NotFoundPage from '~/pages/NotFoundPage.vue';
import OpenImage from '~/pages/OpenImage.vue';

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
    path: '/images/:id',
    name: 'OpenImage',
    component: OpenImage,
  },
  {
    path: '/',
    name: 'Authorized',
    component: AuthLayout,
    children: [
      {
        path: '/',
        name: 'main',
        component: MainPage,
      },
      {
        path: '/folder/:folderId',
        name: 'folder',
        component: FolderPage,
      },
      {
        path: '/generate',
        name: 'generate',
        component: GeneratePage,
      },
      {
        path: '/prints',
        name: 'Prints',
        component: PrintsPage,
      },
      {
        path: '/password',
        name: 'Password',
        component: PasswordPage,
      },
      {
        path: '/:pathMatch(.*)*',
        name: 'NotFound',
        component: NotFoundPage,
      },
    ],
  },
];

const router = createRouter({
  history: createWebHistory(),
  routes,
});

export default router;
