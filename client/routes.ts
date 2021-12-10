// import {createWebHistory, createRouter} from 'vue-router';
// import AuthorizedLayout from './layouts/AuthorizedLayout/AuthorizedLayout.vue';
// import AuthPage from './pages/AuthPage.vue';
// import RegistrationPage from './pages/RegistrationPage.vue';
// import PasswordPage from './pages/PasswordPage.vue';
// import MainPage from './pages/MainPage.vue';
// import EditProjectPage from './pages/EditProjectPage.vue';
// import DetailPage from './pages/DetailPage.vue';
// import ContainersPage from './pages/ContainersPage.vue';
// import VolumesPage from './pages/VolumesPage.vue';
// import StatisticsPage from './pages/StatisticsPage.vue';
// import MembersPage from './pages/MembersPage.vue';
// import NotFoundPage from './pages/NotFoundPage.vue';

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
    component: AuthorizedLayout,
    children: [
      {
        path: '/password',
        name: 'Password',
        component: PasswordPage,
      },
      {
        path: '/',
        name: 'Main',
        component: MainPage,
      },
      {
        path: '/projects/create',
        name: 'CreateProject',
        component: EditProjectPage,
      },
      {
        path: '/projects/:id',
        name: 'ProjectDetail',
        component: DetailPage,
      },
      {
        path: '/projects/:id/edit',
        name: 'EditProject',
        component: EditProjectPage,
      },
      {
        path: '/projects/:id/edit',
        name: 'EditProject',
        component: EditProjectPage,
      },
      {
        path: '/containers',
        name: 'Containers',
        component: ContainersPage,
      },
      {
        path: '/volumes',
        name: 'Volumes',
        component: VolumesPage,
      },
      {
        path: '/server',
        name: 'Statistics',
        component: StatisticsPage,
      },
      {
        path: '/members',
        name: 'Members',
        component: MembersPage,
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
