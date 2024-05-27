import {
  Injectable,
  NotFoundException,
  BadRequestException,
} from '@nestjs/common';
import { InjectModel } from '@nestjs/mongoose';
import { Model, isValidObjectId } from 'mongoose';
import { Service } from './schemas/service.schema';
import { CreateServiceDto } from './dto/create-service.dto';
import { UpdateServiceDto } from './dto/update-service.dto';

@Injectable()
export class ServicesService {
  constructor(
    @InjectModel(Service.name) private readonly serviceModel: Model<Service>,
  ) {}

  async findAll(deleted: boolean): Promise<Service[]> {
    return this.serviceModel.find({ deleted }).exec();
  }

  async findOne(id: string, deleted: boolean): Promise<Service> {
    this.validateObjectId(id);

    const service = await this.serviceModel.findOne({ id, deleted }).exec();
    if (!service) {
      throw new NotFoundException(`Service with ID ${id} not found`);
    }

    return service;
  }

  async findByName(name: string, deleted: boolean): Promise<Service> {
    const service = this.serviceModel.findOne({ name, deleted }).exec();
    if (!service) {
      throw new NotFoundException(`Service with name ${name} not found`);
    }

    return service;
  }

  async create(createServiceDto: CreateServiceDto): Promise<Service> {
    if (await this.findByName(createServiceDto.name, false)) {
      throw new BadRequestException('Service already exists');
    }

    const createdService = new this.serviceModel(createServiceDto);
    return createdService.save();
  }

  async update(
    id: string,
    updateServiceDto: UpdateServiceDto,
  ): Promise<Service> {
    await this.remove(id);
    return this.create(updateServiceDto as CreateServiceDto);
  }

  async remove(id: string): Promise<Service> {
    this.validateObjectId(id);

    const existingService = await this.findOne(id, false);
    if (!existingService) {
      throw new NotFoundException(`Service with Id ${id} not found`);
    }

    existingService.deleted = true;
    return existingService.save();
  }

  private validateObjectId(id: string): void {
    if (!isValidObjectId(id)) {
      throw new BadRequestException(`Invalid Id format: ${id}`);
    }
  }
}
