import {createStore, useStore, Store} from 'vuex';
import {UserType} from '../types/user';

type MainStateTypes = {
    user: UserType,
    isAuthorized: boolean,
}

export default createStore<MainStateTypes>({
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
