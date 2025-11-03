export const isHttpUrl = (value?: string) => {
  if (!value) return false
  try {
    const u = new URL(value)
    return u.protocol === 'http:' || u.protocol === 'https:'
  } catch {
    return false
  }
}

export const isYouTubeUrl = (value?: string) => {
  if (!value) return false
  try {
    const u = new URL(value)
    if (!['http:', 'https:'].includes(u.protocol)) return false
    const host = u.hostname.toLowerCase()
    const allowedHosts = new Set([
      'youtube.com',
      'www.youtube.com',
      'm.youtube.com',
      'youtu.be',
      'www.youtu.be',
    ])
    if (!allowedHosts.has(host)) return false

    const p = u.pathname
    if (host.endsWith('youtu.be')) {
      return p.split('/')[1]?.length > 0
    }
    if (p.startsWith('/watch')) {
      return !!u.searchParams.get('v')
    }
    if (p.startsWith('/shorts/')) {
      return p.split('/')[2]?.length > 0
    }
    if (p.startsWith('/playlist')) {
      return !!u.searchParams.get('list')
    }
    if (
      p.startsWith('/channel/') ||
      p.startsWith('/user/') ||
      p.startsWith('/c/') ||
      p.startsWith('/@')
    ) {
      return true
    }
    return true
  } catch {
    return false
  }
}
