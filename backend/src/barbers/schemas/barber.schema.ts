import { Prop, Schema, SchemaFactory } from '@nestjs/mongoose';
import { Document } from 'mongoose';

@Schema()
export class Barber extends Document {
  @Prop({ required: true })
  name: string;

  @Prop()
  photo: string;

  @Prop({ default: false })
  is_working: boolean;
}

export const BarberSchema = SchemaFactory.createForClass(Barber);
