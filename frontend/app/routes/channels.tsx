import {useState} from 'react'
import {
  Table,
  Button,
  Group,
  LoadingOverlay,
  Title,
  ActionIcon,
  Checkbox,
} from '@mantine/core'
import {useDisclosure} from '@mantine/hooks'
import {IconPencil, IconTrash} from '@tabler/icons-react'
import {useChannels} from '../hooks/useChannels'
import {ChannelModal} from '../components/ChannelModal'
import type {Channel, UpdateChannelDTO} from '../types/channel'

export default function ChannelsPage() {
  const {
    channels,
    loading,
    createChannel,
    updateChannel,
    deleteChannel,
    setOnlyWithImage,
  } = useChannels()
  const [modalOpened, {open: openModal, close: closeModal}] =
    useDisclosure(false)
  const [selectedChannel, setSelectedChannel] = useState<Channel | null>(null)

  const handleAdd = () => {
    setSelectedChannel(null)
    openModal()
  }

  const handleEdit = (channel: Channel) => {
    setSelectedChannel(channel)
    openModal()
  }

  const handleSubmit = (values: UpdateChannelDTO) => {
    if (selectedChannel) {
      updateChannel(selectedChannel.id, values)
    } else {
      createChannel(values)
    }
    closeModal()
  }

  const rows = channels.map(channel => (
    <Table.Tr key={channel.id}>
      <Table.Td>{channel.id}</Table.Td>
      <Table.Td>{channel.nome}</Table.Td>
      <Table.Td>{channel.tag}</Table.Td>
      <Table.Td>{channel.url}</Table.Td>
      <Table.Td>{channel.youtube}</Table.Td>
      <Table.Td>{channel.imagem}</Table.Td>
      <Table.Td>{channel.site}</Table.Td>
      <Table.Td>
        <Group gap="xs">
          <ActionIcon variant="subtle" onClick={() => handleEdit(channel)}>
            <IconPencil size={16} />
          </ActionIcon>
          <ActionIcon
            variant="subtle"
            color="red"
            onClick={() => deleteChannel(channel.id)}
          >
            <IconTrash size={16} />
          </ActionIcon>
        </Group>
      </Table.Td>
    </Table.Tr>
  ))

  return (
    <div style={{padding: '2rem', backgroundColor: '#f5f5f5', height: '100vh'}}>
      <Group justify="space-between" mb="xl">
        <Title order={2}>Channels</Title>
        <Checkbox
          label="Only with Image"
          onChange={event => {
            setOnlyWithImage(event.currentTarget.checked)
          }}
        />
        <Button onClick={handleAdd}>Add Channel</Button>
      </Group>
      <div style={{position: 'relative'}}>
        <LoadingOverlay visible={loading} />
        <Table striped highlightOnHover>
          <Table.Thead>
            <Table.Tr>
              <Table.Th>ID</Table.Th>
              <Table.Th>Name</Table.Th>
              <Table.Th>Tag</Table.Th>
              <Table.Th>URL</Table.Th>
              <Table.Th>YouTube</Table.Th>
              <Table.Th>Image</Table.Th>
              <Table.Th>Site</Table.Th>
              <Table.Th>Actions</Table.Th>
            </Table.Tr>
          </Table.Thead>
          <Table.Tbody>{rows}</Table.Tbody>
        </Table>
      </div>

      <ChannelModal
        opened={modalOpened}
        onClose={closeModal}
        onSubmit={handleSubmit}
        initialValues={
          selectedChannel
            ? {
                nome: selectedChannel.nome,
                url: selectedChannel.url,
                youtube: selectedChannel.youtube,
                tag: selectedChannel.tag,
                imagem: selectedChannel.imagem,
                site: selectedChannel.site,
              }
            : undefined
        }
        title={selectedChannel ? 'Edit Channel' : 'Add Channel'}
      />
    </div>
  )
}
