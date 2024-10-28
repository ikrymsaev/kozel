export abstract class WS {
  protected connection: WebSocket | null = null

  protected conn(url: string) {
    if (this.connection) {
      this.connection.close()
    }
    this.connection = new WebSocket(url)
  }

  protected disconn() {
    if (this.connection) {
      this.connection.close()
    }
    this.connection = null
  }

  protected withConn = <T>(method: (connection: WebSocket) => T) => {
    if (!this.connection) {
      throw new Error('No connection')
    }
    return method(this.connection)
  }

  // private getConnUrl({ roomId, userId, username }: TConnectionParams) {
  //   return `ws://localhost:8080/ws/join/${roomId}?user_id=${userId}&username=${username}`
  // }
}

export type TConnectionParams = {
  lobbyId: string
  userId: string
  username: string
}
