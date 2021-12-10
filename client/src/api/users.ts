import $axios from './index';
import {CreateUserType, UserType} from '~/types/user';

export const signIn = (email: string, password: string) => $axios.post<UserType>('user/sign-in', {
  email,
  password,
});

export const signUp = (user: CreateUserType) => $axios.post<UserType>('user/sign-up', user);

export const logout = () => $axios.post<boolean>('user/logout');

export const refresh = () => $axios.post<boolean>('user/refresh');

export const changePassword = () => $axios.post<boolean>('user/password');

export const getUser = () => $axios.post<UserType>('user');

export const changeUser = (user: Partial<UserType>) => $axios.post<UserType>('user/password', user);
