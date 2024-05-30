import {
  BadRequestException,
  Injectable,
  NotFoundException,
} from '@nestjs/common';
import { InjectModel } from '@nestjs/mongoose';
import { Model } from 'mongoose';
import { Barber } from './schemas/barber.schema';
import { CreateBarberDto } from './dto/create-barber.dto';
import { UpdateBarberDto } from './dto/update-barber.dto';
import { validateObjectId } from '../utils/objectId.util';

@Injectable()
export class BarbersService {
  constructor(@InjectModel(Barber.name) private barberModel: Model<Barber>) {}

  async create(createBarberDto: CreateBarberDto): Promise<Barber> {
    if (await this.findByName(createBarberDto.name, false)) {
      throw new BadRequestException('Barber already exists');
    }

    const newBarber = new this.barberModel(createBarberDto);
    return newBarber.save();
  }

  async findAll(deleted: boolean): Promise<Barber[]> {
    return this.barberModel.find({ deleted }).exec();
  }

  async findByName(name: string, deleted: boolean): Promise<Barber> {
    const barber = await this.barberModel.findOne({ name, deleted }).exec();
    if (!barber) {
      throw new NotFoundException(`Barber with name ${name} not found`);
    }

    return barber;
  }

  async findOne(id: string, deleted: boolean): Promise<Barber> {
    validateObjectId(id);

    const barber = await this.barberModel.findOne({ id, deleted }).exec();
    if (!barber) {
      throw new NotFoundException(`Barber with Id ${id} not found`);
    }

    return barber;
  }

  async update(id: string, updateBarberDto: UpdateBarberDto): Promise<Barber> {
    validateObjectId(id);

    const barber = await this.findByName(updateBarberDto.name, false);
    if (barber && barber.id !== id) {
      throw new BadRequestException('Name already exists in another barber');
    }

    const existingBarber = await this.findOne(id, false);
    if (!existingBarber) {
      throw new NotFoundException(`Barber with Id ${id} not found`);
    }

    return this.barberModel
      .findByIdAndUpdate(id, updateBarberDto, { new: true })
      .exec();
  }

  async remove(id: string): Promise<Barber> {
    validateObjectId(id);

    const existingBarber = await this.findOne(id, false);
    if (!existingBarber) {
      throw new NotFoundException(`Barber with Id ${id} not found`);
    }

    existingBarber.deleted = true;
    return existingBarber.save();
  }
}
