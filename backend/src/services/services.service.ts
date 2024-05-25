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

  async findAll(): Promise<Service[]> {
    return this.serviceModel.find().exec();
  }

  async findOne(id: string): Promise<Service> {
    this.validateObjectId(id);

    const service = await this.serviceModel.findById(id).exec();
    if (!service) {
      throw new NotFoundException(`Service with ID ${id} not found`);
    }

    return service;
  }

  async findByName(name: string): Promise<Service> {
    const service = this.serviceModel.findOne({ name }).exec();
    if (!service) {
      throw new NotFoundException(`Service with name ${name} not found`);
    }

    return service;
  }

  async create(createServiceDto: CreateServiceDto): Promise<Service> {
    if (await this.findByName(createServiceDto.name)) {
      throw new BadRequestException('Service already exists');
    }

    const createdService = new this.serviceModel(createServiceDto);
    return createdService.save();
  }

  async update(
    id: string,
    updateServiceDto: UpdateServiceDto,
  ): Promise<Service> {
    this.validateObjectId(id);

    const service = await this.findByName(updateServiceDto.name);
    if (service && service.id !== id) {
      throw new BadRequestException('Service already exists');
    }

    const existingService = await this.serviceModel.findById(id).exec();
    if (!existingService) {
      throw new NotFoundException(`Service with Id ${id} not found`);
    }

    return this.serviceModel
      .findByIdAndUpdate(id, updateServiceDto, { new: true })
      .exec();
  }

  async remove(id: string): Promise<Service> {
    this.validateObjectId(id);

    const existingService = await this.serviceModel.findById(id).exec();
    if (!existingService) {
      throw new NotFoundException(`Service with Id ${id} not found`);
    }

    return this.serviceModel.findByIdAndDelete(id).exec();
  }

  private validateObjectId(id: string): void {
    if (!isValidObjectId(id)) {
      throw new BadRequestException(`Invalid Id format: ${id}`);
    }
  }
}
