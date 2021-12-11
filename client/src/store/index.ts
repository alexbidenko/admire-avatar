import {InjectionKey} from 'vue';
import {createStore, useStore, Store} from 'vuex';
import {UserType} from '~/types/user';
import {ImageType} from '~/types/image';

type MainState = {
  user: UserType,
  avatar: ImageType | null;
  isAuthorized: boolean,
}

export const key: InjectionKey<Store<MainState>> = Symbol();

export default createStore<MainState>({
  state: () => {
    return {
      user: {} as UserType,
      avatar: null,
      isAuthorized: false,
    };
  },
  mutations: {
    setUser: (state, value: UserType) => {
      state.user = value;
      state.isAuthorized = true;
    },
    setAvatar: (state, value: ImageType) => {
      state.avatar = value;
    },
    logout: (state) => {
      state.user = {} as UserType;
      state.isAuthorized = false;
    },
  },
});

export const useMainStore = () => useStore(key);
