import { Prop, Schema, SchemaFactory } from '@nestjs/mongoose';
import { Document } from 'mongoose';
import { ApiProperty } from '@nestjs/swagger';

@Schema()
export class Service extends Document {
  @ApiProperty({ description: 'Id of the service' })
  _id: string;

  @Prop({ required: true })
  @ApiProperty({ description: 'Name of the service' })
  name: string;

  @Prop()
  @ApiProperty({ description: 'Description of the service' })
  description?: string;

  @Prop({ default: 0 })
  @ApiProperty({ description: 'Duration of the service in minutes' })
  duration: number;

  @Prop({ type: [String], enum: ['haircut', 'beard', 'eyebrows', 'other'] })
  @ApiProperty({ description: 'Kinds of the service', type: [String] })
  kinds: string[];

  @Prop({ default: false })
  @ApiProperty({ description: 'Service is a combo?' })
  is_combo: boolean;

  @Prop({ required: true })
  @ApiProperty({ description: 'Price of the service' })
  price: number;

  @Prop()
  @ApiProperty({ description: 'Commission rate for the service' })
  commission_rate?: number;
}

export const ServiceSchema = SchemaFactory.createForClass(Service);
