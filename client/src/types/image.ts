/* eslint-disable no-unused-vars */
enum ImageTypeEnum {
  Avatar = 'avatar',
  Print = 'print',
}
/* eslint-enable no-unused-vars */

export type BaseImageType = {
  source: string;
};

export type ImageType = {
  id: number;
  userId: number;
  createdAt: string;
  updatedAt: string;
  type: ImageTypeEnum;
} & BaseImageType;

export type ImageInputType = {
  phrase: string;
  tags: string[];
};