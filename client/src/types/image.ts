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
