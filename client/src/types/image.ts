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
  folderId: number;
  createdAt: string;
  updatedAt: string;
  type: ImageTypeEnum;
  main: boolean;
} & BaseImageType;

export type ImageInputType = {
  phrase: string;
  tags: string[];
};

export type FolderType = {
  id: number;
  name: string;
  userId: number;
  createdAt: string;
  updatedAt: string;
  public: boolean;
};
