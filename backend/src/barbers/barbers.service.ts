import {
  BadRequestException,
  Injectable,
  NotFoundException,
} from '@nestjs/common';
import { InjectModel } from '@nestjs/mongoose';
import { isValidObjectId, Model } from 'mongoose';
import { Barber } from './schemas/barber.schema';
import { CreateBarberDto } from './dto/create-barber.dto';
import { UpdateBarberDto } from './dto/update-barber.dto';

@Injectable()
export class BarbersService {
  constructor(@InjectModel(Barber.name) private barberModel: Model<Barber>) {}

  async create(createBarberDto: CreateBarberDto): Promise<Barber> {
    if (await this.findByName(createBarberDto.name)) {
      throw new BadRequestException('Barber already exists');
    }

    const newBarber = new this.barberModel(createBarberDto);
    return newBarber.save();
  }

  async findAll(): Promise<Barber[]> {
    return this.barberModel.find().exec();
  }

  async findByName(name: string): Promise<Barber> {
    const barber = await this.barberModel.findOne({ name }).exec();
    if (!barber) {
      throw new NotFoundException(`Barber with name ${name} not found`);
    }

    return barber;
  }

  async findOne(id: string): Promise<Barber> {
    this.validateObjectId(id);

    const barber = await this.barberModel.findById(id).exec();
    if (!barber) {
      throw new NotFoundException(`Barber with Id ${id} not found`);
    }

    return barber;
  }

  async update(id: string, updateBarberDto: UpdateBarberDto): Promise<Barber> {
    this.validateObjectId(id);

    const barber = await this.findByName(updateBarberDto.name);
    if (barber && barber.id !== id) {
      throw new BadRequestException('Name already exists in another barber');
    }

    const existingBarber = await this.barberModel.findById(id).exec();
    if (!existingBarber) {
      throw new NotFoundException(`Barber with Id ${id} not found`);
    }

    return this.barberModel
      .findByIdAndUpdate(id, updateBarberDto, { new: true })
      .exec();
  }

  async remove(id: string): Promise<Barber> {
    this.validateObjectId(id);

    const existingBarber = await this.barberModel.findById(id).exec();
    if (!existingBarber) {
      throw new NotFoundException(`Barber with Id ${id} not found`);
    }

    return this.barberModel.findByIdAndDelete(id).exec();
  }

  private validateObjectId(id: string): void {
    if (!isValidObjectId(id)) {
      throw new BadRequestException(`Invalid Id format: ${id}`);
    }
  }
}
