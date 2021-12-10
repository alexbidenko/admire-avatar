import $axios from './index';
import {BaseImageType, ImageInputType, ImageType} from '../src/types/user';

export const getImages = () => $axios.get<ImageType[]>('images');

export const generateImage = (data: ImageInputType) => $axios.put<BaseImageType>('images', data);

export const saveImage = (data: BaseImageType) => $axios.post<ImageType>('images', data);

export const deleteImage = (imageId: number) => $axios.delete<boolean>(`images/${imageId}`);

export const createAvatar = (imageId: number) => $axios.put<boolean>(`images/${imageId}`);
