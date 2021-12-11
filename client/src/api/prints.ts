import $axios from '~/api/index';
import {ImageInputType, ImageType} from '~/types/image';

export const getPrints = () => $axios.get<ImageType[]>('prints/0/100');

export const generatePrints = (data: { count: number } & ImageInputType) => $axios.post('prints', data);

export const saveAsAvatar = (imageId: number) => $axios.put(`prints/${imageId}`);

export const clear = () => $axios.delete('prints');
