export interface Channel {
  id: number
  nome: string
  url: string
  youtube: string
  tag: string
  imagem: string | null
  site: string
}

export type UpdateChannelDTO = Partial<
  Omit<Channel, 'ID' | 'CreatedAt' | 'UpdatedAt' | 'DeletedAt'>
>

export interface ChannelInfo {
  id: string
  url: string
  youtube: string
  tag: string
  imagem: string | null
  site: string
}

export type ChannelMap = Record<string, ChannelInfo>
