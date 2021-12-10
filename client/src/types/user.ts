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
