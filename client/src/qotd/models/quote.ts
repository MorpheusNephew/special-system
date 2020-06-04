export interface ErrorResponse {
  message: string;
  code: number;
}

export interface QuoteResponse {
  author: string;
  language: string;
  likes: number;
  quote: string;
  tags: string[];
}

export interface Response {
  data?: QuoteResponse;
  error?: ErrorResponse;
}
