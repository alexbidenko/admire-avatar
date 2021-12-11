import {InjectionKey} from 'vue';
import {createStore, useStore, Store} from 'vuex';
import {UserType} from '~/types/user';

type MainState = {
  user: UserType,
  isAuthorized: boolean,
}

export const key: InjectionKey<Store<MainState>> = Symbol();

export default createStore<MainState>({
  state: () => {
    return {
      user: {} as UserType,
      isAuthorized: false,
    };
  },
  mutations: {
    setUser: (state, value: UserType) => {
      state.user = value;
      state.isAuthorized = true;
    },
    logout: (state) => {
      state.user = {} as UserType;
      state.isAuthorized = false;
    },
  },
});

export const useMainStore = () => useStore(key);
