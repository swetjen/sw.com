export interface IProjectItem {
  logo: string
  brand: string
  hero_image: string
  title: string
  titleExtended?: string
  subtitle?: string
  summary: string
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
