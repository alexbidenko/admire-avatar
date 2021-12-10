import {createApp} from 'vue';
import App from './App.vue';
import router from './routes';
import store, {key} from '~/store';

const app = createApp(App);

app.use(router).use(store, key).mount('#app');
