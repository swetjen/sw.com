export interface IProjectItem {
  logo: string
  brand: string
  heroImage: string
  dates: string
  title: string
  titleExtended?: string
  subtitle?: string
  slug: string
}

export interface IContactForm {
  name: string
  message: string
  email: string
}

export interface IContactFormResponse {
  status: string
  guid: string
}