export interface Channel {
    ID: number;
    CreatedAt: string;
    UpdatedAt: string;
    DeletedAt: string | null;
    nome: string;
    url: string;
    youtube: string;
    tag: string;
    imagem: string | null;
    site: string;
  }
  
  export type UpdateChannelDTO = Partial<Omit<Channel, 'ID' | 'CreatedAt' | 'UpdatedAt' | 'DeletedAt'>>;
  