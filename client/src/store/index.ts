import {InjectionKey} from 'vue';
import {createStore, useStore, Store} from 'vuex';
import {UserType} from '~/types/user';

type MainState = {
  user: UserType,
  isAuthorized: boolean,
  avatar: number;
}

export const key: InjectionKey<Store<MainState>> = Symbol();

export default createStore<MainState>({
  state: () => ({
    user: {} as UserType,
    isAuthorized: false,
    avatar: 1,
  }),
  mutations: {
    setUser: (state, value: UserType) => {
      state.user = value;
      state.isAuthorized = true;
    },
    logout: (state) => {
      state.user = {} as UserType;
      state.isAuthorized = false;
    },
    setAvatar: (state, value: number) => {
      state.avatar = value;
    },
  },
});

export const useMainStore = () => useStore(key);
