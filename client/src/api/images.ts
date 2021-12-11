import $axios from './index';
import {BaseImageType, ImageInputType, ImageType} from '~/types/image';

export const getImages = () => $axios.get<ImageType[]>('images/0/100');

export const generateImage = (data: ImageInputType) => $axios.put<BaseImageType>('images', data);

export const saveImage = (data: BaseImageType) => $axios.post<ImageType>('images', data);

export const deleteImage = (imageId: number) => $axios.delete<boolean>(`images/${imageId}`);

export const createAvatar = (imageId: number) => $axios.put<boolean>(`images/${imageId}`);

export const getAvatar = () => $axios.get<ImageType>('images/avatar');
