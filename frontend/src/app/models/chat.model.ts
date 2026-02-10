export interface Message {
  role: string;
  content: string;
}

export interface ChatRequest {
  prompt: string;
}

export interface GroqRequest {
  model: string;
  messages: Message[];
  temperature: number;
  max_tokens: number;
}

export interface GroqResponse {
  choices: {
    message: Message;
  }[];
}
