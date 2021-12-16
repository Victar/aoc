package adventofcode.year2021;

import java.util.ArrayList;
import java.util.List;

import org.apache.commons.lang3.StringUtils;
import org.junit.Test;

import adventofcode.BaseTest;
import lombok.Data;

public class Day16 extends BaseTest {

	@Test public void runSingle() throws Exception {
		final String input = "CE00C43D881120";
		final Package pack = new Package(convertToBin(input));
		System.out.println(pack.getValue());
	}

	@Test public void runSilver() throws Exception {
		final ArrayList<String> data = readStringFromFile("year2021/day16/input.txt");
		final Package pack = new Package(convertToBin(data.get(0)));
		System.out.println(pack.getSumVersion());
	}

	@Test public void runGold() throws Exception {
		final ArrayList<String> data = readStringFromFile("year2021/day16/input.txt");
		final Package pack = new Package(convertToBin(data.get(0)));
		System.out.println(pack.getValue());
	}

	public String convertToBin(final String input) {
		final StringBuilder sb = new StringBuilder();
		for (int i = 0; i < input.length(); i++) {
			final int inputInt = Integer.valueOf("" + input.charAt(i), 16);
			String curStr = Integer.toString(inputInt, 2);
			while (curStr.length() != 4) {
				curStr = "0" + curStr;
			}
			sb.append(curStr);
		}
		return sb.toString();
	}

	@Data static class Package {

		String input;
		int current;
		int version;
		int type;
		Integer typeID;
		String binValue = "";
		List<Package> childPackages = new ArrayList<>();

		public Package(final String input) {
			this.input = input;
			this.version = readNextInt(3);
			this.type = readNextInt(3);

			if (this.type != 4) {
				this.typeID = readNextInt(1); //I
				if (this.typeID == 0) {
					final int subPackageLengthInBits = readNextInt(15); // L = 27
					String subPackagesString = readNext(subPackageLengthInBits);
					boolean addChildren = true;
					while (addChildren) {
						final Package child = new Package(subPackagesString);
						this.childPackages.add(child);
						final int childCurrent = child.getCurrent();
						subPackagesString = subPackagesString.substring(childCurrent);
						addChildren = StringUtils.isNotEmpty(subPackagesString);
					}
				} else if (this.typeID == 1) {
					final int subPackageNumber = readNextInt(11);
					String subPackagesString = this.input.substring(this.current);
					for (int i = 0; i < subPackageNumber; i++) {
						final Package child = new Package(subPackagesString);
						this.childPackages.add(child);
						final int childCurrent = child.getCurrent();
						this.current += childCurrent;
						subPackagesString = subPackagesString.substring(childCurrent);
					}
				}
			} else {
				boolean keepReading = true;
				while (keepReading) {
					final String current = readNext(5);
					keepReading = current.startsWith("1");
					this.binValue += current.substring(1, 5);
				}
			}
		}

		public long getValue() {
			if (this.type == 4) {
				return Long.parseLong(this.binValue, 2);
			}
			if (this.type == 0) {
				long result = 0l;
				for (final Package p : this.childPackages) {
					result += p.getValue();
				}
				return result;
			}
			if (this.type == 1) {
				long result = 1l;
				for (final Package p : this.childPackages) {
					result = result * p.getValue();
				}
				return result;
			}
			if (this.type == 2) {
				long result = Long.MAX_VALUE;
				for (final Package p : this.childPackages) {
					result = Math.min(result, p.getValue());
				}
				return result;
			}
			if (this.type == 3) {
				long result = Long.MIN_VALUE;
				for (final Package p : this.childPackages) {
					result = Math.max(result, p.getValue());
				}
				return result;
			}
			if (this.type == 5) {
				return this.childPackages.get(0).getValue() > this.childPackages.get(1).getValue() ? 1l : 0l;
			}
			if (this.type == 6) {
				return this.childPackages.get(0).getValue() < this.childPackages.get(1).getValue() ? 1l : 0l;
			}
			if (this.type == 7) {
				return this.childPackages.get(0).getValue() == this.childPackages.get(1).getValue() ? 1l : 0l;
			}
			return 0l;
		}

		public int getSumVersion() {
			int result = this.version;
			for (final Package p : this.childPackages) {
				result += p.getSumVersion();
			}
			return result;
		}

		public Integer readNextInt(final int size) {
			final String result = this.input.substring(this.current, this.current + size);
			this.current += size;
			return Integer.parseInt(result, 2);
		}

		public String readNext(final int size) {
			final String result = this.input.substring(this.current, this.current + size);
			this.current += size;
			return result;
		}
	}

}
