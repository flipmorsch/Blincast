import {useState} from 'react'
import {Modal, Button, TextInput, Group, Box} from '@mantine/core'
import {useForm} from '@mantine/form'
import type {UpdateChannelDTO} from '../types/channel'

interface ChannelModalProps {
  opened: boolean
  onClose: () => void
  onSubmit: (values: UpdateChannelDTO) => void
  initialValues?: UpdateChannelDTO
  title: string
}

export function ChannelModal({
  opened,
  onClose,
  onSubmit,
  initialValues,
  title,
}: ChannelModalProps) {
  const form = useForm<UpdateChannelDTO>({
    initialValues: initialValues || {
      nome: '',
      url: '',
      youtube: '',
      tag: '',
      imagem: '',
      site: '',
    },
    validate: {
      nome: value => (value && value.length > 0 ? null : 'Name is required'),
      url: value =>
        value && /^(ftp|http|https):\/\/[^ "]+$/.test(value)
          ? null
          : 'A valid URL is required',
      site: value =>
        value && /^(ftp|http|https):\/\/[^ "]+$/.test(value)
          ? null
          : 'A valid URL is required',
      tag: value => (value && value.length > 0 ? null : 'Tag is required'),
    },
  })

  return (
    <Modal opened={opened} onClose={onClose} title={title}>
      <Box component="form" onSubmit={form.onSubmit(onSubmit)}>
        <TextInput
          label="Name"
          placeholder="Channel Name"
          defaultValue={initialValues?.nome}
          required
        />
        <TextInput
          mt="sm"
          label="URL"
          placeholder="https://example.com"
          defaultValue={initialValues?.url}
          required
        />
        <TextInput
          mt="sm"
          label="YouTube"
          placeholder="https://youtube.com/channel/..."
          defaultValue={initialValues?.youtube}
        />
        <TextInput
          mt="sm"
          label="Tag"
          placeholder="Technology"
          defaultValue={initialValues?.tag}
          required
        />
        <TextInput
          mt="sm"
          label="Image URL"
          placeholder="https://example.com/image.png"
          defaultValue={initialValues?.imagem!}
        />
        <TextInput
          mt="sm"
          label="Site"
          placeholder="https://example.com"
          defaultValue={initialValues?.site}
          required
        />
        <Group justify="flex-end" mt="md">
          <Button variant="default" onClick={onClose}>
            Cancel
          </Button>
          <Button type="submit">Submit</Button>
        </Group>
      </Box>
    </Modal>
  )
}
