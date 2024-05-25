import { Injectable } from '@nestjs/common';
import { InjectModel } from '@nestjs/mongoose';
import { Model } from 'mongoose';
import { Barber } from './schemas/barber.schema';
import { CreateBarberDto } from './dto/create-barber.dto';
import { UpdateBarberDto } from './dto/update-barber.dto';

@Injectable()
export class BarbersService {
  constructor(@InjectModel(Barber.name) private barberModel: Model<Barber>) {}

  async create(createBarberDto: CreateBarberDto): Promise<Barber> {
    const newBarber = new this.barberModel(createBarberDto);
    return newBarber.save();
  }

  async findAll(): Promise<Barber[]> {
    return this.barberModel.find().exec();
  }

  async findByName(name: string): Promise<Barber> {
    return this.barberModel.findOne({ name }).exec();
  }

  async findOne(id: string): Promise<Barber> {
    return this.barberModel.findById(id).exec();
  }

  async update(id: string, updateBarberDto: UpdateBarberDto): Promise<Barber> {
    return this.barberModel
      .findByIdAndUpdate(id, updateBarberDto, { new: true })
      .exec();
  }

  async remove(id: string): Promise<Barber> {
    return this.barberModel.findByIdAndDelete(id).exec();
  }
}
