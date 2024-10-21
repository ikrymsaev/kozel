export abstract class WS {
  protected connection: WebSocket | null = null

  protected conn(params: TConnectionParams) {
    if (this.connection) {
      this.connection.close()
    }
    this.connection = new WebSocket(this.getConnUrl(params))
  }

  protected withConn = <T>(method: (connection: WebSocket) => T) => {
    if (!this.connection) {
      throw new Error('No connection')
    }
    return method(this.connection)
  }

  private getConnUrl({ roomId, userId, username }: TConnectionParams) {
    return `ws://localhost:8080/ws/join/${roomId}?user_id=${userId}&username=${username}`
  }
}

export type TConnectionParams = {
  roomId: string
  userId: string
  username: string
}
