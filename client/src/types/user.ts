export type CreateUserType = {
  name: string;
  email: string;
  password: string;
};

export type UserType = {
  id: number;
  createdAt: string;
  updatedAt: string;
} & Omit<CreateUserType, 'password'>;

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
