export type UserType = {
  id: number;
  name: string;
  email: string;
  createdAt: string;
  updatedAt: string;
  password?: string;
};

export type BaseImageType = {
  source: string;
};

export type ImageType = {
  id: string;
  userId: number;
  createdAt: string;
  updatedAt: string;
} & BaseImageType;

export type ImageInputType = {
  phrase: string;
  tags: string[];
};
