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
}

export type TConnectionParams = {
  lobbyId: string
  userId: string
  username: string
}
