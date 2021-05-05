export class TokenStorage {
  public static getToken(): string | null {
    return sessionStorage.getItem("APIKEY");
  }

  public static setToken(token: string): void {
    sessionStorage.setItem("APIKEY", token);
  }
  public static removeToken(): void {
    sessionStorage.removeItem("APIKEY");
  }
}
