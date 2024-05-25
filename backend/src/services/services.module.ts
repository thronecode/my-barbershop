import { Module } from '@nestjs/common';
import { ServicesService } from './services.service';
import { ServicesController } from './services.controller';

@Module({
  providers: [ServicesService],
  controllers: [ServicesController],
})
export class ServicesModule {}
