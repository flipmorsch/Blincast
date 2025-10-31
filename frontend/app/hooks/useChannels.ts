import {useEffect, useState} from 'react'
import {notifications} from '@mantine/notifications'
import type {Channel, UpdateChannelDTO, ChannelMap} from '../types/channel'

const API_URL = 'http://localhost:8080'

const removeEmptyFields = <T extends object>(obj: T): Partial<T> => {
  const newObj: Partial<T> = {}
  for (const key in obj) {
    if (obj[key] !== '' && obj[key] !== null) {
      newObj[key] = obj[key]
    }
  }
  return newObj
}

export function useChannels() {
  const [channels, setChannels] = useState<Channel[]>([])
  const [loading, setLoading] = useState(true)
  const [onlyWithImage, setOnlyWithImage] = useState(false)

  const fetchChannels = async () => {
    setLoading(true)
    try {
      const fetchUrl = onlyWithImage
        ? `${API_URL}/channel/with-image`
        : `${API_URL}/channel`
      const response = await fetch(fetchUrl)
      if (!response.ok) {
        throw new Error('Failed to fetch channels')
      }
      const data = (await response.json()) as Record<string, Partial<Channel>>
      const mappedData = Object.entries(data).map(([nome, channelObj]) => ({
        nome,
        ...(channelObj || {}),
      }))
      setChannels(mappedData as unknown as Channel[])
    } catch (error) {
      notifications.show({
        title: 'Error',
        message:
          error instanceof Error ? error.message : 'An unknown error occurred',
        color: 'red',
      })
    } finally {
      setLoading(false)
    }
  }

  useEffect(() => {
    fetchChannels()
  }, [onlyWithImage])

  const createChannel = async (channel: Omit<UpdateChannelDTO, 'ID'>) => {
    try {
      const payload = removeEmptyFields(channel)
      const response = await fetch(`${API_URL}/channel/`, {
        method: 'POST',
        headers: {'Content-Type': 'application/json'},
        body: JSON.stringify(payload),
      })
      if (!response.ok) {
        throw new Error('Failed to create channel')
      }
      notifications.show({
        title: 'Success',
        message: 'Channel created successfully',
        color: 'green',
      })
      fetchChannels() // Refresh list
    } catch (error) {
      notifications.show({
        title: 'Error',
        message:
          error instanceof Error ? error.message : 'An unknown error occurred',
        color: 'red',
      })
    }
  }

  const updateChannel = async (id: number, channel: UpdateChannelDTO) => {
    try {
      const payload = removeEmptyFields(channel)
      const response = await fetch(`${API_URL}/channel/${id}`, {
        method: 'PUT',
        headers: {'Content-Type': 'application/json'},
        body: JSON.stringify(payload),
      })
      if (!response.ok) {
        throw new Error('Failed to update channel')
      }
      notifications.show({
        title: 'Success',
        message: 'Channel updated successfully',
        color: 'green',
      })
      fetchChannels() // Refresh list
    } catch (error) {
      notifications.show({
        title: 'Error',
        message:
          error instanceof Error ? error.message : 'An unknown error occurred',
        color: 'red',
      })
    }
  }

  const deleteChannel = async (id: number) => {
    try {
      const response = await fetch(`${API_URL}/channel/${id}`, {
        method: 'DELETE',
      })
      if (!response.ok) {
        throw new Error('Failed to delete channel')
      }
      notifications.show({
        title: 'Success',
        message: 'Channel deleted successfully',
        color: 'green',
      })
      fetchChannels() // Refresh list
    } catch (error) {
      notifications.show({
        title: 'Error',
        message:
          error instanceof Error ? error.message : 'An unknown error occurred',
        color: 'red',
      })
    }
  }

  return {
    channels,
    loading,
    createChannel,
    updateChannel,
    deleteChannel,
    onlyWithImage,
    setOnlyWithImage,
  }
}
