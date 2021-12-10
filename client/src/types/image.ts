export type BaseImageType = {
  source: string;
};

export type ImageType = {
  id: number;
  userId: number;
  createdAt: string;
  updatedAt: string;
} & BaseImageType;

export type ImageInputType = {
  phrase: string;
  tags: string[];
};
