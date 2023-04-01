export interface IProjectItem {
  logo: string
  brand: string
  heroImage: string
  title: string
  titleExtended?: string
  slug: string
}

export interface IContactForm {
  message: string
  email: string
}

export interface INotification {
  msg: string,
  class?: string
}
