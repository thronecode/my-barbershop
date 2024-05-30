import { Module } from '@nestjs/common';
import { MongooseModule } from '@nestjs/mongoose';
import { BarbersService } from './barbers.service';
import { BarbersController } from './barbers.controller';
import { Barber, BarberSchema } from './schemas/barber.schema';

@Module({
  imports: [
    MongooseModule.forFeature([{ name: Barber.name, schema: BarberSchema }]),
  ],
  controllers: [BarbersController],
  providers: [BarbersService],
})
export class BarbersModule {}
